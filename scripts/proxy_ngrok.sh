#!/bin/bash

# Check if ngrok is installed globally
if command -v ngrok &> /dev/null
then
    NGROK_CMD="ngrok"
else
    # Check if ngrok is available in the local path
    if [ -f "./ngrok" ]
    then
        NGROK_CMD="./ngrok"
    else
        echo "ngrok could not be found. Please install ngrok to proceed."
        exit 1
    fi
fi

# Run ngrok to proxy TCP traffic on port 8080
$NGROK_CMD tcp 8080