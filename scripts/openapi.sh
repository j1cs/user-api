#!/usr/bin/env bash
set -e

RED="\033[31m"
WHITE="\033[37m"
RESET="\033[0m"
BOLD=$(tput bold)
NORMAL=$(tput sgr0)

readonly SERVICE="$1"
readonly OUTPUT_DIR="$2"
readonly PACKAGE="$3"

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

if ! command -v oapi-codegen &> /dev/null
then
    echo -e "${RED}${BOLD}Error: ensure you have go bin path configured${NORMAL}${RESET}"
    echo "You should have this in your .bashrc or .zshrc:"
    echo "export GOPATH=\$HOME/go\nexport GOBIN=\$GOPATH/bin\nexport PATH=\$PATH:\$GOBIN"
    echo "Copy and paste in your shell configuration (.bashrc or .zshrc)"
    exit 1
fi

MODEL_TEMP=$(mktemp)
SERVER_TEMP=$(mktemp)
CLIENT_TEMP=$(mktemp)

sed "s|%package%|$PACKAGE|g" scripts/templates/model.tmpl > "$MODEL_TEMP"
sed "s|%package%|$PACKAGE|g" scripts/templates/server.tmpl > "$SERVER_TEMP"
sed "s|%package%|$PACKAGE|g" scripts/templates/client.tmpl > "$CLIENT_TEMP"

oapi-codegen --config "$MODEL_TEMP" "api/$SERVICE.yaml" > "$OUTPUT_DIR/${SERVICE}_types.gen.go"
oapi-codegen --config "$SERVER_TEMP" "api/$SERVICE.yaml" > "$OUTPUT_DIR/${SERVICE}_server.gen.go"
oapi-codegen --config "$CLIENT_TEMP" "api/$SERVICE.yaml" > "$OUTPUT_DIR/${SERVICE}_client.gen.go"

echo -e "${WHITE}${BOLD}Code generated successfully${NORMAL}${RESET}"
rm "$MODEL_TEMP" "$SERVER_TEMP" "$CLIENT_TEMP"
