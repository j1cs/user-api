#!/usr/bin/env bash

set -e

RED="\033[31m"
GREEN="\033[32m"
RESET="\033[0m"
BOLD=$(tput bold)
NORMAL=$(tput sgr0)

readonly EXEC="${1:-dev}"

go install github.com/cespare/reflex@latest
go install github.com/go-delve/delve/cmd/dlv@latest

if ! command -v reflex &> /dev/null; then
    echo -e "${RED}${BOLD}Error: ensure you have go bin path configured${NORMAL}${RESET}"
    echo "You should have this in your .bashrc or .zshrc"
    echo "export GOPATH=\$HOME/go\nexport GOBIN=\$GOPATH/bin\nexport PATH=\$PATH:\$GOBIN"
    echo "Copy and paste in your shell configuration (.bashrc or .zshrc)"
    exit 1
fi

if ! command -v curl &> /dev/null; then
  echo -e "${RED}curl could not be found"
  echo "Please install curl on your machine https://curl.se/"
  echo "Then run 'make $EXEC' again${RESET}"
  exit 1
fi

trap "make down" SIGINT

#we should check some services before start the api.
echo -e "${GREEN}Waiting for service to start...${RESET}"
while [ -z $(curl -s http://localhost:$PUBSUB_EMULATOR_PORT) ]; do
    printf '.'
    sleep 1
done

until ./scripts/database_health.sh; do
    printf '.'
    sleep 1
done

echo -e "${GREEN}Service is up!${RESET}"

export PUBSUB_EMULATOR_HOST=localhost:$PUBSUB_EMULATOR_PORT

if [ "$EXEC" == "dev" ]; then
  reflex -r '(\.go$|go\.mod)' -s go run ./cmd/api/main.go
else
  reflex -r '(\.go$|go\.mod)' -s -- sh -c ./scripts/dlv.sh
fi
