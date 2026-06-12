#!/bin/bash

set -e

echo "🚀 Building fakestack v2.0.0 for GitHub Release..."

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

VERSION="2.0.0"

for platform in "${platforms[@]}"; do
    GOOS="${platform%/*}"
    GOARCH="${platform#*/}"
    
    output="../bin/fakestack-${GOOS}-${GOARCH}"
    if [ "$GOOS" = "windows" ]; then
        output="${output}.exe"
    fi
    
    echo "📦 Building for $GOOS/$GOARCH..."
    
    # Only enable CGO for macOS (SQLite on macOS works well)
    if [ "$GOOS" = "darwin" ]; then
        export CGO_ENABLED=1
    else
        # Disable CGO for cross-platform builds to avoid toolchain issues
        export CGO_ENABLED=0
    fi
    
    GOOS=$GOOS GOARCH=$GOARCH go build \
        -ldflags="-s -w -X main.Version=$VERSION" \
        -o "$output"
    
    # Make executable on Unix
    if [ "$GOOS" != "windows" ]; then
        chmod +x "$output"
    fi
    
    # Display size
    ls -lh "$output"
done

echo ""
echo "✅ Build complete!"
echo ""
echo "📊 Binary sizes:"
ls -lh ../bin/ | tail -n +2
echo ""
echo "Total size: $(du -sh ../bin | cut -f1)"
echo ""
echo "Ready to create GitHub release v$VERSION"
