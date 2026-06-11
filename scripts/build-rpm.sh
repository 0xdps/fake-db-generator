#!/bin/bash
# Build .rpm package for RedHat/Fedora/CentOS

set -e

VERSION="${1:-1.2.0}"
ARCH="${2:-x86_64}"

echo "Building fakestack v${VERSION} .rpm package for ${ARCH}..."

# Create RPM build structure
RPMBUILD_DIR="$HOME/rpmbuild"
mkdir -p "$RPMBUILD_DIR"/{BUILD,RPMS,SOURCES,SPECS,SRPMS}

# Create spec file
cat > "$RPMBUILD_DIR/SPECS/fakestack.spec" <<EOF
Name:           fakestack
Version:        ${VERSION}
Release:        1%{?dist}
Summary:        High-performance database generator with realistic fake data

License:        MIT
URL:            https://github.com/0xdps/fake-stack
Source0:        fakestack-linux-${ARCH}

%description
Generate database tables and populate them with realistic fake data from JSON schemas.
Supports SQLite, PostgreSQL, MySQL, MariaDB, MS SQL Server, and CockroachDB.
Features 116+ data generators and 10-50x faster than pure Python/JavaScript implementations.

%prep

%build

%install
mkdir -p %{buildroot}%{_bindir}
install -m 755 %{SOURCE0} %{buildroot}%{_bindir}/fakestack

%files
%{_bindir}/fakestack

%changelog
* $(date "+%a %b %d %Y") 0xdps <your-email@example.com> - ${VERSION}-1
- Initial package
EOF

# Copy binary to SOURCES
cp "../bin/fakestack-linux-${ARCH}" "$RPMBUILD_DIR/SOURCES/"

# Build RPM
rpmbuild -ba "$RPMBUILD_DIR/SPECS/fakestack.spec"

echo "✓ Package built: $RPMBUILD_DIR/RPMS/${ARCH}/fakestack-${VERSION}-1.${ARCH}.rpm"
