# 1. Build stage
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o readiness_probe main.go

# 2. Final stage
FROM alpine:latest

RUN adduser -D nonrootuser

USER nonrootuser

EXPOSE 8000

COPY --from=builder /app/readiness_probe /usr/local/bin/readiness_probe

CMD ["readiness_probe"]
