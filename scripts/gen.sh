#!/bin/bash

# Define the URL of the workflow endpoint
URL="http://localhost:8082/wf-gen"

# Make the request and store the response
response=$(curl -s -X GET "$URL")

# Check if the response is not empty
if [ -n "$response" ]; then
    echo "Workflow JSON Response:"
    echo "$response" | jq .  # Pretty print with jq, if installed
else
    echo "Failed to retrieve workflow. Ensure the server is running."
fi
