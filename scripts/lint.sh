#!/usr/bin/env bash

# ANSI escape sequences for setting text colors and styles
RED="\033[31m"
WHITE="\033[37m"
RESET="\033[0m"
BOLD=$(tput bold)
NORMAL=$(tput sgr0)

go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

if ! command -v golangci-lint &> /dev/null
then
    echo -e "${RED}${BOLD}Error: ensure you have go bin path configured${NORMAL}${RESET}"
    echo "You should have this in your .bashrc or .zshrc:"
    echo "export GOPATH=\$HOME/go\nexport GOBIN=\$GOPATH/bin\nexport PATH=\$PATH:\$GOBIN"
    echo "Copy and paste in your shell configuration (.bashrc or .zshrc)"
    exit 1
fi

golangci-lint run internal/... | tee /dev/stderr
