#!/bin/bash
# Build .deb package for Debian/Ubuntu

set -e

VERSION="${1:-1.2.0}"
ARCH="${2:-amd64}"

echo "Building fakestack v${VERSION} .deb package for ${ARCH}..."

# Create package structure
PKG_DIR="fakestack_${VERSION}_${ARCH}"
mkdir -p "$PKG_DIR/DEBIAN"
mkdir -p "$PKG_DIR/usr/local/bin"

# Copy binary
cp "../bin/fakestack-linux-${ARCH}" "$PKG_DIR/usr/local/bin/fakestack"
chmod +x "$PKG_DIR/usr/local/bin/fakestack"

# Create control file
cat > "$PKG_DIR/DEBIAN/control" <<EOF
Package: fakestack
Version: ${VERSION}
Section: database
Priority: optional
Architecture: ${ARCH}
Maintainer: 0xdps <your-email@example.com>
Description: High-performance database generator with realistic fake data
 Generate database tables and populate them with realistic fake data from JSON schemas.
 Supports SQLite, PostgreSQL, MySQL, MariaDB, MS SQL Server, and CockroachDB.
 Features 116+ data generators and 10-50x faster than pure Python/JavaScript implementations.
Homepage: https://github.com/0xdps/fake-stack
EOF

# Build package
dpkg-deb --build "$PKG_DIR"

echo "✓ Package built: ${PKG_DIR}.deb"
