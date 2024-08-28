#!/bin/bash

# Check if go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go to proceed."
    exit 1
fi

# Set the Go module
export GO111MODULE=on

# Run the Go application
go run main.go
