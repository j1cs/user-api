#!/usr/bin/env bash

# ANSI escape sequences for setting text colors and styles
RED="\033[0;31m"
WHITE="\033[0;37m"
NOCOLOR="\033[0m"
BOLD=$(tput bold)
NORMAL=$(tput sgr0)

# Install goimports
GOIMPORTS=$(go install golang.org/x/tools/cmd/goimports@latest 2>&1)

# Check if the command failed
if [ $? -ne 0 ]; then
    echo -e "${RED}${BOLD}Failed to install goimports:${NORMAL}${NOCOLOR}"
    echo -e "${RED}$GOIMPORTS${NOCOLOR}"
    exit 1
fi

# Check if goimports is installed
if ! command -v goimports &> /dev/null; then
    echo -e "${RED}${BOLD}Error: goimports not found! Ensure you have Go bin path configured.${NORMAL}${NOCOLOR}"
    echo "You should have this in your .bashrc or .zshrc:"
    echo -e "${WHITE}export GOPATH=\$HOME/go\nexport GOBIN=\$GOPATH/bin\nexport PATH=\$PATH:\$GOBIN${NOCOLOR}"
    echo "Copy and paste in your shell configuration (.bashrc or .zshrc)"
    exit 1
fi

# Format your Go code
GOFMT=$(gofmt -l -w internal/ 2>&1)

# Check if the command failed
if [ $? -ne 0 ]; then
    echo -e "${RED}${BOLD}Failed to format your Go code with gofmt:${NORMAL}${NOCOLOR}"
    echo -e "${RED}$GOFMT${NOCOLOR}"
    exit 1
fi

# Use goimports on your Go code
GOIMPORTS=$(goimports -l -w internal/ 2>&1)

# Check if the command failed
if [ $? -ne 0 ]; then
    echo -e "${RED}${BOLD}Failed to format your Go code with goimports:${NORMAL}${NOCOLOR}"
    echo -e "${RED}$GOIMPORTS${NOCOLOR}"
    exit 1
fi

echo -e "${GREEN}${BOLD}Successfully formatted your Go code!${NORMAL}${NOCOLOR}"
