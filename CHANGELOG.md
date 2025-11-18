# Changelog

All notable changes to Fakestack will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **Template Generator**: Custom generator for creating data patterns using templates
  - Syntax: `{{generator|modifier1|modifier2}}`
  - Supports all built-in generators (name, email, word, uuid, random_int, etc.)
  - Modifiers: `upper`, `lower`, `title`, `trim`, `truncate(n)`
  - Example: `"{{word|upper|truncate(3)}}-{{random_int(1000,9999)}}"` → `PRD-4821`
  - Perfect for SKUs, employee IDs, license plates, order numbers
  - See `docs/TEMPLATE_EXAMPLES.md` for comprehensive examples
  - Enables custom data formats without writing code
- **Interactive Schema Generator**: New `-g` / `--generate` CLI flag for creating schemas
  - Built directly into fakestack CLI (no separate script needed)
  - 10 pre-built templates (Users, Employees, Products, Orders, Customers, Blog Posts, Inventory, Transactions, Students, Tasks)
  - Support for all 6 database types
  - Customizable database credentials and output filename
  - Automatic row count suggestions per template
  - Usage: `fakestack -g .` or `fakestack -g my-schema.json`
  - Implementation in `golang/templates.go` for clean code organization
- **New Generators**: Added 80+ generators across 12 categories for comprehensive test data generation
  - **Financial/Payment** (8): `credit_card`, `credit_card_type`, `credit_card_cvv`, `credit_card_exp`, `currency`, `currency_long`, `bitcoin_address`, `bitcoin_private_key`
  - **Time & Timestamps** (7): `timestamp`, `time`, `year`, `month`, `month_string`, `weekday`, `timezone`
  - **Localization** (4): `country_code`, `language`, `language_abbr`, `locale`
  - **Product/E-commerce** (8): `color`, `hex_color`, `safe_color`, `product_name`, `product_category`, `product_description`, `product_feature`, `price`
  - **Files & Media** (4): `filename`, `file_extension`, `mime_type`, `image_url`
  - **User Agent & Browser** (5): `user_agent`, `chrome_user_agent`, `firefox_user_agent`, `safari_user_agent`, `opera_user_agent`
  - **Books & Media** (5): `book_title`, `book_author`, `book_genre`, `movie_name`, `movie_genre`
  - **Animals & Nature** (7): `animal`, `animal_type`, `pet_name`, `cat`, `dog`, `bird`, `farm_animal`
  - **Food** (8): `fruit`, `vegetable`, `breakfast`, `lunch`, `dinner`, `snack`, `dessert`, `drink`
  - **Vehicle/Transportation** (5): `car_maker`, `car_model`, `car_type`, `car_fuel_type`, `car_transmission_type`
  - **Identifiers** (4): `ssn`, `ein`, `iban`, `routing_number`
  - **Text Types** (6): `emoji`, `emoji_description`, `emoji_category`, `quote`, `phrase`, `question`
  - **App & Software** (3): `app_name`, `app_version`, `app_author`
  - **Range Support**: Added `integer`, `float`, and `decimal` generators with min/max range support
    - `integer` / `random_int` - Generate integers with configurable min/max range
    - `float` / `decimal` - Generate floating-point numbers with configurable min/max range
    - Backward compatible with existing `random_int` generator
  - Total generators: 116+ covering most common data generation needs
- **Database Support**: Added support for 3 additional database systems
  - MariaDB - MySQL-compatible database with enhanced features
  - MS SQL Server - Microsoft's enterprise database with IDENTITY syntax support
  - CockroachDB - Distributed SQL database (PostgreSQL-compatible)
- **Manual Testing Workflow**: Created `.github/workflows/test-databases.yml` for on-demand database testing
  - Tests all 6 supported databases in parallel
  - Manual trigger via GitHub Actions UI (`workflow_dispatch`)
  - Individual checkboxes for selective database testing (all selected by default)
  - Docker-based testing for MySQL, PostgreSQL, MariaDB, MSSQL, SQLite, CockroachDB
- **Documentation**: Added comprehensive configuration examples and Docker setup for new databases
- **Version Selector**: Added version dropdown to documentation (powered by mike)
- **Uninstall Script**: Created `scripts/uninstall.sh` for removing fakestack from npm, pip, and Homebrew

### Changed
- **Homebrew Tap**: Renamed from `homebrew-fakestack` to `homebrew-packages` for future tool support
  - New installation: `brew tap 0xdps/packages && brew install fakestack`
  - Prepares tap for additional tools beyond fakestack
- **Database Driver**: Added Microsoft SQL Server driver (`github.com/denisenkom/go-mssqldb`)
- **Documentation**: Updated all installation instructions to use new tap name

### Fixed
- **MSSQL Compatibility**: Fixed SQL syntax issues for MS SQL Server
  - Added square brackets `[table]` and `[column]` around identifiers
  - Changed placeholders from `?` to `@p1, @p2, @p3` format
  - Updated `CREATE TABLE` to use `IF OBJECT_ID` syntax
  - Updated `DROP TABLE` to use `IF OBJECT_ID` syntax
  - Fixed sqlcmd path in workflow to `/opt/mssql-tools18/bin/sqlcmd` with `-C` flag
- **Deprecation Fix**: Replaced deprecated `strings.Title()` with Unicode-compliant `cases.Title()`
  - Added dependency: `golang.org/x/text v0.31.0`
  - Fixed deprecation warning in Go 1.24.0
  - Improved Unicode handling for title casing

### Technical Details
- Total supported databases: **6** (SQLite, MySQL, PostgreSQL, MariaDB, MSSQL, CockroachDB)
- MariaDB uses MySQL driver (drop-in compatible)
- CockroachDB uses PostgreSQL driver
- MSSQL implements IDENTITY autoincrement and NEWID() for random selection

## [1.0.1] - 2025-11-17

### Added
- Added version flag (`-v`, `--version`) to Go binary to display version information
- Version is now injected at build time using ldflags

### Fixed
- Fixed logo display across all package READMEs and documentation
  - Python README: Uses absolute GitHub URL for PyPI compatibility
  - Node.js README: Uses relative path (works on npm)
  - Documentation: Uses symlinked assets folder
- Fixed binary build paths in publish workflow
- Fixed test assertions to use dynamic version checking instead of hardcoded values

### Changed
- Improved package documentation visibility on PyPI and npm
- Updated all READMEs with consistent branding and logo display
- Enhanced version consistency validation in CI/CD workflows

## [1.0.0] - 2025-11-17

### 🚀 Initial Release - High-Performance Database Generator

First stable release of Fakestack - a high-performance database generator powered by Go with multi-language support.

### Added

#### Go Core Engine ⚡
- Complete rewrite in Go for **10-50x performance improvement**
  - `golang/main.go` - CLI entry point with -c, -p, -f, -d flags
  - `golang/schema.go` - JSON schema parsing and data structures
  - `golang/database.go` - Multi-database support (SQLite, MySQL, PostgreSQL)
  - `golang/generator.go` - 50+ fake data generators using gofakeit
  - `golang/populate.go` - Data population with progress tracking

#### Cross-Platform Build System 📦
- `scripts/build.sh` - Automated build script for all platforms
- Support for Linux (amd64, arm64), macOS (amd64, arm64), Windows (amd64)
- Binary size optimization with stripped symbols (~45MB total for all platforms)

#### Python Wrapper 🐍
- New `python/fakestack/runner.py` - Lightweight wrapper with platform detection
- External process execution pattern (subprocess)
- Zero runtime dependencies (all dependencies removed)
- 100% backward compatible CLI interface
- Published as `fakestack` on PyPI

#### npm Package 📦
- New `node/` package for Node.js ecosystem
- Full TypeScript support with type definitions
- CLI and programmatic API
- Platform detection and binary execution
- Published as `fakestack` on npm
- CommonJS module format for maximum compatibility

#### CI/CD Pipeline 🔄
- Multi-language testing (Go + Python + Node.js)
- Automated testing on Python 3.8-3.13 and Node.js 18-22
- Cross-platform binary builds
- Dual publishing to PyPI and npm

#### Documentation 📚
- Language-specific READMEs: `golang/README.md`, `python/README.md`, `node/README.md`
- Updated root documentation for multi-language architecture
- Updated `CONTRIBUTING.md` for new structure
- LICENSE moved to root directory

### Changed

- **Repository Structure** - Language-based organization:
  - `golang/` - Go core implementation
  - `python/` - Python wrapper (was: root-level `fakestack/`)
  - `node/` - Node.js wrapper
  - Configuration files moved to respective language folders

- **Package Naming** - Unified across ecosystems:
  - PyPI: `fakestack` (was: `fakestack_py`)
  - npm: `fakestack` (was: `@fakestack/core`)
  - Import: `from fakestack import ...` (Python), `require('fakestack')` (Node.js)

- **Removed all legacy Python implementation code**
  - Removed `models/fake.py` (Faker providers)
  - Removed `models/schema.py` (Pydantic models)
  - Removed `models/utils.py` (utility functions)
  - Removed `data/` directory (example schemas)
  
- **Updated dependencies**
  - Removed: `faker`, `SQLAlchemy`, `pydantic`, `psycopg2-binary`, `mysql-connector-python`
  - Python package now has **zero runtime dependencies**
  - Node.js package has **zero runtime dependencies**
  - All functionality provided by bundled Go binaries

- **Updated CI/CD workflows**
  - Paths updated for `golang/`, `python/`, `node/` structure
  - Separate test jobs for each language
  - Build binaries once, test across all platforms

- **Updated documentation**
  - README.md reflects multi-language architecture
  - Repository structure section added
  - Development setup for Go, Python, and Node.js
  - Acknowledgments updated to reflect Go core (gofakeit)

### Technical Details

- **Architecture**: External process execution (subprocess pattern)
- **Go version**: 1.21+
- **Python support**: 3.8-3.13
- **Node.js support**: 18+
- **Platforms**: Linux, macOS, Windows (amd64, arm64)
- **Go Dependencies**: gofakeit/v7, go-sqlite3, go-mysql-driver, lib/pq
- **Module Format**: CommonJS for Node.js (maximum compatibility)

### Performance

| Dataset | Python v2.0 | Go Core v2.1 | Speedup |
|---------|-------------|--------------|---------|
| 1K rows | ~2.5s | ~0.1s | **25x** ⚡ |
| 10K rows | ~25s | ~0.8s | **31x** ⚡ |
| 100K rows | ~250s | ~6s | **42x** ⚡ |

### Breaking Changes

**For Python Users**: None! The CLI interface remains identical:
```bash
fakestack -d .
fakestack -c -p -f schema.json
```

**For Contributors**: New directory structure requires path updates:
- Python code: Now in `python/` folder
- Node.js code: Now in `node/` folder
- Go core: Now in `golang/` folder
- See [CONTRIBUTING.md](CONTRIBUTING.md) for updated setup

### Migration Guide

**For Python Users**: No action required. Install or upgrade as usual:
```bash
pip install --upgrade fakestack
```

**For Node.js Users**: Install via npm:
```bash
npm install fakestack
```

**For Contributors**: Update paths in development workflow:
```bash
# Python development
cd python && pip install -e ".[dev]"

# Node.js development
cd node && npm install

# Go development
cd golang && go build
```

### Installation

**Python:**
```bash
pip install fakestack
```

**Node.js:**
```bash
npm install fakestack
```

**Homebrew:**
```bash
brew tap 0xdps/packages
brew install fakestack
```

### Quick Start

```bash
# Download example schema
fakestack -d .

# Create tables and populate data
fakestack -c -p -f schema.json
```

---

**Full Documentation**: https://fake-stack.readthedocs.io/

[Unreleased]: https://github.com/0xdps/fake-stack/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/0xdps/fake-stack/releases/tag/v1.0.0
