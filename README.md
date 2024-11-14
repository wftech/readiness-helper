
# Readiness Probe Sidecar

This repository contains a small Go application used as a readiness probe, 
serving HTTP responses to signal readiness and termination status.

## Features
- Serves a `200 OK` response on `/healthz` by default.
- Responds with a `410 GONE` status on `/healthz` upon receiving a `SIGTERM` signal or a manual shutdown request at `/stop`.
- Configurable port via command-line parameter.

## Requirements
- Docker
- Go 1.23+ (if building manually)

## Usage

### Build and Run with Docker

1. **Build the Docker Image**

   ```bash
   docker build -t readiness-probe .
   ```

2. **Run the Docker Container**

   ```bash
   docker run -p 8000:8000 readiness-probe
   ```

   The application will listen on port 8000 by default. You can map this port or change it using `-p`.

### Configuration Options

- `-p` : Sets the port for the server (default: 8000).
- `-h` : Displays help for available flags.

Example to run on a custom port:

```bash
readiness_probe -p 8080
```

## Dockerfile Explanation

The Dockerfile is multi-stage:

1. **Build Stage** - Compiles the Go application.
2. **Final Stage** - Copies the built binary into a minimal Alpine image, creating a lightweight container.

The application runs as a non-root user for added security.

## Endpoints

- **/healthz**: Returns `200 OK` if healthy, `410 GONE` if marked as terminated.
- **/stop**: Sets the application to return `410 GONE` for `/healthz`, emulating a graceful shutdown.

## Example Docker Commands

Retrieve the container ID and stop it with SIGTERM:

```bash
container_id=$(docker ps -qf "ancestor=readiness-probe")
docker kill -s SIGTERM "$container_id"
```

