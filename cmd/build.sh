#!/bin/bash

# go build
go build -ldflags "-w -s -v -X main.VERSION=1.0.0 -X 'main.BUILD_TIME=$(date)' -X 'main.GO_VERSION=$(go version)'"
