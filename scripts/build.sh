#!/bin/sh

# compile the server using the cgo
CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o release/point-server ./cmd/server