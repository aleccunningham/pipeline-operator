#!/bin/sh

export CGO_ENABLED=0
export GOOS=linux

go build -ldflags '-extldflags "-static"' -a -v -o bin/linux/duke-crd ./cmd/duke-crd
