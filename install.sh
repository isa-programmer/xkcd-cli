#!/bin/bash

set -e

# Detect platform and architecture
OS="$(uname -s)"
ARCH="$(uname -m)"

# Default values
BINARY_URL=""

if [[ "$OS" == "Linux" ]]; then
    if [[ "$ARCH" == "x86_64" || "$ARCH" == "amd64" ]]; then
        BINARY_URL="https://github.com/isa-programmer/xkcd-cli/releases/latest/download/xkcd-linux-amd64"
    else
        echo "Unsupported Linux architecture: $ARCH"
        exit 1
    fi
elif [[ "$OS" == "Darwin" ]]; then
    if [[ "$ARCH" == "arm64" ]]; then
        # Apple Silicon (M1/M2)
        BINARY_URL="https://github.com/isa-programmer/xkcd-cli/releases/latest/download/xkcd-darwin-arm64"
    elif [[ "$ARCH" == "x86_64" ]]; then
        # Apple Intel
        BINARY_URL="https://github.com/isa-programmer/xkcd-cli/releases/latest/download/xkcd-darwin-amd64"
    else
        echo "Unsupported macOS architecture: $ARCH"
        exit 1
    fi
else
    echo "Unsupported OS: $OS"
    exit 1
fi

echo "Downloading xkcd-cli for $OS/$ARCH..."
sudo wget -O /usr/local/bin/xkcd "$BINARY_URL"
sudo chmod +x /usr/local/bin/xkcd

echo "xkcd-cli has been installed successfully!"
