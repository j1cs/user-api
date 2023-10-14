#!/usr/bin/env bash
set -e

if ! command -v mockery &> /dev/null; then
    echo "mockery could not be found"
    echo "Installing from https://vektra.github.io/mockery/"
    echo "We recommend using homebrew or your favorite package manager"
    exit 1
fi

interfaces=("UserRepository" "UserService" "UserPublisher")

pwd

for interface in "${interfaces[@]}"; do
    mockery --name="$interface" --output="./internal/v1/domain/mocks" --dir="./internal/v1/domain" --recursive
done
