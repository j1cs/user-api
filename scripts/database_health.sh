#!/usr/bin/env bash

host="localhost"
port="5432"
timeout=30
retry_interval=5

if ! command -v nc &> /dev/null; then
    echo "'nc' (Netcat) is not installed. Please install it for your operating system and try again."
    exit 1
fi

while true; do
  echo "Checking connection to $host:$port with timeout $timeout seconds..."

  if nc -z -w $timeout $host $port; then
    echo "Port is ready and accepting connections..."
    exit 0
  else
    echo "Connection error: Unable to connect to $host:$port. Retrying in $retry_interval seconds..."
    sleep $retry_interval
  fi
done