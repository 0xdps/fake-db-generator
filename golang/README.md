# Fakestack Go Core

High-performance core implementation of Fakestack in Go, providing the engine for fake data generation and database population. This binary is used by both the Python and Node.js wrappers.

## Prerequisites

- **Go 1.21+** - [Download](https://go.dev/dl/)

## Quick Start

```bash
# Install dependencies
go mod download

# Build binary
go build -o fakestack

# Run
./fakestack -d .
./fakestack -cpf schema.json
```

## Building

**Single platform:**
```bash
go build -o fakestack
```

**Cross-compile all platforms:**
```bash
../scripts/build.sh
```

This creates binaries in `../bin/`:
- `fakestack-linux-amd64`
- `fakestack-linux-arm64`
- `fakestack-darwin-amd64` (Intel Mac)
- `fakestack-darwin-arm64` (Apple Silicon)
- `fakestack-windows-amd64.exe`

**Manual cross-compilation:**
```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ../bin/fakestack-linux-amd64
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ../bin/fakestack-darwin-arm64
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ../bin/fakestack-windows-amd64.exe
```

## Usage

```bash
# Download example schema
./fakestack -d <directory>

# Create tables from schema
./fakestack -c -f schema.json

# Populate tables with fake data
./fakestack -p -f schema.json

# Create and populate in one command
./fakestack -cpf schema.json
```

## Features

✅ **50+ Fake Data Generators** - Names, emails, addresses, phone numbers, dates, etc.  
✅ **Multi-Database Support** - SQLite, MySQL, PostgreSQL  
✅ **Foreign Key Relationships** - Automatic FK resolution  
✅ **Unique Value Generation** - Guaranteed unique values for unique columns  
✅ **Custom Generators** - person, user, random_from, etc.  
✅ **Progress Tracking** - Real-time feedback on data generation  
✅ **High Performance** - 10-50x faster than pure Python implementations

## Performance

| Dataset | Python v2.0 | Go Core | Speedup |
|---------|-------------|---------|---------|
| 1K rows | ~2.5s | ~0.1s | **25x** ⚡ |
| 10K rows | ~25s | ~0.8s | **31x** ⚡ |
| 100K rows | ~250s | ~6s | **42x** ⚡ |

## Architecture

```
golang/
├── main.go       # CLI entry point, flag parsing
├── schema.go     # JSON schema parsing and validation
├── database.go   # Database connection and operations
├── generator.go  # Fake data generation using gofakeit
├── populate.go   # Data population with FK resolution
└── go.mod        # Go dependencies
```

## Dependencies

- **gofakeit/v7** - Comprehensive fake data generation
- **go-sqlite3** - SQLite driver (CGO)
- **go-mysql-driver** - MySQL driver
- **lib/pq** - PostgreSQL driver

## Development

**Run tests:**
```bash
go test -v ./...
go test -race ./...
go test -cover ./...
```

**Format code:**
```bash
go fmt ./...
```

**Lint:**
```bash
golangci-lint run
```

## Integration

This Go binary is called by:
- **Python wrapper** (`python/fakestack/runner.py`) via subprocess
- **Node.js wrapper** (`node/src/index.ts`) via child_process

Wrappers detect platform, select appropriate binary, and execute with same CLI flags.
