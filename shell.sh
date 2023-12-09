#!/bin/bash

# Set the host and port of the echo server
ECHO_SERVER_HOST="localhost"
ECHO_SERVER_PORT="7878"

# Command to run (e.g., sending "Hello, World!")
COMMAND="Hello, World!"

# Number of times to run the command
LOOP_COUNT=1000000

# Loop to connect to the echo server and run the command
for ((i=1; i<=$LOOP_COUNT; i++)); do
    # Use netcat to connect to the echo server and send the command
    echo "$COMMAND" | nc -N "$ECHO_SERVER_HOST" "$ECHO_SERVER_PORT"

    # Sleep for a short duration between each iteration (adjust as needed)
    # sleep 1
done
