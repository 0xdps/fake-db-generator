#!/bin/bash

set -e

echo "🚀 Building fakestack for all platforms..."

cd "$(dirname "$0")/../golang"

# Clean previous builds
rm -rf ../bin
mkdir -p ../bin

# Build for all platforms
platforms=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

for platform in "${platforms[@]}"; do
    GOOS="${platform%/*}"
    GOARCH="${platform#*/}"
    
    output="../bin/fakestack-${GOOS}-${GOARCH}"
    if [ "$GOOS" = "windows" ]; then
        output="${output}.exe"
    fi
    
    echo "📦 Building for $GOOS/$GOARCH..."
    
    # CGO is required for SQLite
    export CGO_ENABLED=1
    
    # Set cross-compilation tools
    if [ "$GOOS" = "linux" ] && [ "$GOARCH" = "arm64" ]; then
        export CC=aarch64-linux-gnu-gcc
    elif [ "$GOOS" = "windows" ]; then
        export CC=x86_64-w64-mingw32-gcc
    elif [ "$GOOS" = "darwin" ] && [ "$GOARCH" != "$(go env GOHOSTARCH)" ]; then
        # Cross-compiling for macOS requires OSX SDK
        echo "⚠️  Warning: Cross-compiling for macOS may require additional setup"
    fi
    
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$output"
    
    # Make executable on Unix
    if [ "$GOOS" != "windows" ]; then
        chmod +x "$output"
    fi
done

echo ""
echo "✅ Build complete! Binaries available in bin/"
echo ""
ls -lh ../bin/
echo ""
echo "Total size: $(du -sh ../bin | cut -f1)"
