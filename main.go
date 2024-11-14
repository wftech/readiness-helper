package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

var (
	healthy int32 = 1 // 1 = 200 OK, 0 = 410 Gone
	port    = flag.Int("p", 8000, "port to listen on") // port as flag
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 OK"))
	} else {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("410 GONE"))
	}
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	atomic.StoreInt32(&healthy, 0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Switched to 410 GONE mode"))
}

func main() {
	// Parsing command-line flags
	flag.Parse()

	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/stop", stopHandler)

	// Signal handling for SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		<-sigs
		atomic.StoreInt32(&healthy, 0)
	}()

	address := fmt.Sprintf(":%d", *port)
	fmt.Printf("Starting server on port %d...\n", *port)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
