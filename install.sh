#!/bin/bash
# Fakestack installer - detects OS/arch and installs the right binary

set -e

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
    darwin) OS="darwin" ;;
    linux) OS="linux" ;;
    mingw*|msys*|cygwin*) OS="windows" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

VERSION="${FAKESTACK_VERSION:-latest}"
if [ "$VERSION" = "latest" ]; then
    VERSION=$(curl -s https://api.github.com/repos/0xdps/fake-stack/releases/latest | grep '"tag_name"' | sed -E 's/.*"v([^"]+)".*/\1/')
fi

echo "Installing fakestack v${VERSION} for ${OS}-${ARCH}..."

BINARY="fakestack-${OS}-${ARCH}"
[ "$OS" = "windows" ] && BINARY="${BINARY}.exe"

URL="https://github.com/0xdps/fake-stack/releases/download/v${VERSION}/${BINARY}"

# Download binary
TMP_DIR=$(mktemp -d)
TMP_FILE="${TMP_DIR}/fakestack"

echo "Downloading from ${URL}..."
if command -v curl >/dev/null 2>&1; then
    curl -L -o "$TMP_FILE" "$URL"
elif command -v wget >/dev/null 2>&1; then
    wget -O "$TMP_FILE" "$URL"
else
    echo "Error: curl or wget required"
    exit 1
fi

chmod +x "$TMP_FILE"

# Install to /usr/local/bin or ~/bin
INSTALL_DIR="/usr/local/bin"
if [ ! -w "$INSTALL_DIR" ]; then
    echo "No write permission to $INSTALL_DIR, using sudo..."
    sudo mv "$TMP_FILE" "${INSTALL_DIR}/fakestack"
else
    mv "$TMP_FILE" "${INSTALL_DIR}/fakestack"
fi

rm -rf "$TMP_DIR"

echo "✓ fakestack installed successfully!"
echo ""
echo "Run 'fakestack -g .' to get started"
