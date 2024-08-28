#!/bin/bash

# Check if go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go to proceed."
    exit 1
fi

# Set the Go module
export GO111MODULE=on

# Build the Go application
go build -o build/dist/__go_build_server.exe billyaddlers/go-relaychat/internal/relaychat/server

# Run the Go application
./build/dist/__go_build_server.exe