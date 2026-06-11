#!/bin/bash

set -e

echo "🔧 Installing fakestack locally (with foreign_key fix)..."

# Build the binary
cd golang
go build -o ../bin/fakestack-darwin-arm64
cd ..

# Install to /usr/local/bin
sudo cp bin/fakestack-darwin-arm64 /usr/local/bin/fakestack
sudo chmod +x /usr/local/bin/fakestack

echo "✓ fakestack installed successfully to /usr/local/bin/fakestack"
echo ""
echo "This will override your Homebrew installation."
echo "To revert to Homebrew version, run: brew reinstall fakestack"
