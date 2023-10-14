#!/usr/bin/env bash
set -e

while [ -z $(curl -s http://$PUBSUB_EMULATOR_HOST) ]; do
  sleep 1
done

echo "Starting... please wait"

if [ "$USE_DLV" == "true" ]; then
  dlv --listen=:2345 --headless=true --api-version=2 --log --accept-multiclient exec ./command
else
  ./command
fi
