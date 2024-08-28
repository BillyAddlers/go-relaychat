#!/bin/bash

# Check if go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go to proceed."
    exit 1
fi

echo "Building for Windows 64-bit"

# Set the Go module
export GO111MODULE=on

# Set the build configuration
export GOARCH=amd64
export GOOS=windows

echo "Build is configured for $GOOS $GOARCH"

# Set the build path
build_path="build/dist/go_relaychat_amd64_windows.exe"

# Build the Go application
go build -o $build_path billyaddlers/go-relaychat/internal/relaychat/client

echo "Build is complete. The executable is located at $build_path"
