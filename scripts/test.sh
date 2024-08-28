#!/bin/bash

# Check if go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go to proceed."
    exit 1
fi

# Uncomment if you straight to testing the client
#export RELAYCHAT_URI=0.tcp.ap.ngrok.io:10016

# Set the Go module
export GO111MODULE=on

# Build the Go application
go build -o build/dist/__go_build_client.exe billyaddlers/go-relaychat/internal/relaychat/client

# Run the Go application
./build/dist/__go_build_client.exe