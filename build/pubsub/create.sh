#!/usr/bin/env bash
while [ -z $(curl -s http://$PUBSUB_EMULATOR_HOST) ]; do
  sleep 1
done
echo "creating topic ${2} on ${1}"
python publisher.py ${1} create ${2}
exit 0