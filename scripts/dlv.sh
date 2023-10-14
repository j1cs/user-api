#!/usr/bin/env bash

# ANSI escape sequences for setting text colors and styles
RED="\033[0;31m"
NOCOLOR="\033[0m"
BOLD=$(tput bold)
NORMAL=$(tput sgr0)

# Print the message
printf "%b%b***You should attach your debugger now***%b%b\n" "${RED}" "${BOLD}" "${NORMAL}" "${NOCOLOR}"

# Start the debugger
dlv debug --listen=:2345 --headless=true --api-version=2 --log --accept-multiclient ./cmd/api/main.go
