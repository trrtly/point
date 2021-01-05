#!/bin/sh

echo "building docker images for ${GOOS}/${GOARCH} ..."

# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/point-server ./cmd/server
