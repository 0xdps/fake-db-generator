# Fakestack

[![PyPI](https://img.shields.io/pypi/v/fakestack)](https://pypi.org/project/fakestack/)
[![npm](https://img.shields.io/npm/v/fakestack)](https://www.npmjs.com/package/fakestack)
[![Python](https://img.shields.io/pypi/pyversions/fakestack)](https://pypi.org/project/fakestack/)
[![Node](https://img.shields.io/node/v/fakestack)](https://www.npmjs.com/package/fakestack)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xdps/fake-stack)](https://goreportcard.com/report/github.com/0xdps/fake-stack)

> **High-performance database generator with realistic fake data**

Generate databases from JSON schemas with realistic fake data. **10-50x faster** than pure Python/JavaScript implementations thanks to a Go core with zero-dependency wrappers for Python and Node.js.

## ✨ Features

- 🚀 **Schema-Driven** - Define tables and data in simple JSON format
- ⚡ **High Performance** - Go core delivers 10-50x speed improvement
- 💡 **Realistic Data** - 50+ generators for names, emails, addresses, dates, and more
- 🗄️ **Multi-Database** - Works with SQLite, MySQL, and PostgreSQL
- 🎯 **Simple API** - Easy CLI and programmatic usage
- 🌍 **Cross-Platform** - Linux, macOS, Windows (amd64 & arm64)
- 📦 **Multi-Ecosystem** - Available on PyPI, npm, and Homebrew
- 🔧 **Zero Dependencies** - Self-contained with bundled binaries

## 📦 Installation

Choose your preferred package manager:

**Python (pip)**
```bash
pip install fakestack
```

**Node.js (npm)**
```bash
npm install fakestack
```

**Homebrew (macOS/Linux)**
```bash
brew install 0xdps/fakestack
```

**Go (from source)**
```bash
cd golang && go build
```

**Direct Download**  
Pre-built binaries available on [GitHub Releases](https://github.com/0xdps/fake-stack/releases)

## 🚀 Quick Start

### 1. Download Example Schema

```bash
fakestack -d .
```

This creates a `schema.json` file in the current directory.

### 2. Create Tables and Populate Data

```bash
# All in one command
fakestack -c -p -f schema.json

# Or separately
fakestack -c -f schema.json  # Create tables
fakestack -p -f schema.json  # Populate data
```

### 3. View Your Data

```bash
sqlite3 test.db "SELECT * FROM users LIMIT 5;"
```

## 💻 Usage

### Command Line

```bash
fakestack [OPTIONS]

Options:
  -c, --create-table     Create database tables from schema
  -p, --populate-data    Populate tables with fake data
  -f, --file <path>      Path to JSON schema file
  -d, --download-schema  Download example schema
  -h, --help            Display help message
```

### Python API

```python
from fakestack import fakestack

# Generate database
exit_code = fakestack(['-c', '-p', '-f', 'schema.json'])

# Or use run_fakestack
from fakestack import run_fakestack
run_fakestack(['-d', '.'])  # Download schema
run_fakestack(['-c', '-p', '-f', 'schema.json'])  # Generate
```

### Node.js / TypeScript API

```javascript
// CommonJS
const { fakestack } = require('fakestack');

// ES Modules
import { fakestack } from 'fakestack';

// Generate database
await fakestack(['-c', '-p', '-f', 'schema.json']);

// TypeScript with options
import { fakestack, FakestackOptions } from 'fakestack';

const options: FakestackOptions = {
  createTables: true,
  populateData: true,
  schemaFile: 'schema.json'
};

await fakestack(options);
```

## 📋 Schema Format

Create a `schema.json` file defining your database structure:

```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "test.db"
  },
  "tables": [
    {
      "name": "users",
      "columns": [
        {
          "name": "id",
          "type": "integer",
          "options": {"primary_key": true, "autoincrement": true}
        },
        {
          "name": "username",
          "type": {"name": "string", "args": {"length": 50}},
          "options": {"nullable": false, "unique": true}
        },
        {
          "name": "email",
          "type": {"name": "string", "args": {"length": 100}},
          "options": {"nullable": false, "unique": true}
        },
        {
          "name": "created_at",
          "type": "datetime",
          "options": {"nullable": false}
        }
      ]
    }
  ],
  "populate": [
    {
      "name": "users",
      "count": 100,
      "fields": [
        {"name": "username", "generator": "user_name"},
        {"name": "email", "generator": "email"},
        {"name": "created_at", "generator": "past_date"}
      ]
    }
  ]
}
```

## 🎨 Available Generators

### Personal Data
- `first_name`, `last_name`, `name`
- `email`, `user_name`, `password`
- `phone_number`, `ssn`

### Address
- `address`, `street_address`
- `city`, `state`, `country`
- `postcode`, `latitude`, `longitude`

### Company
- `company`, `company_suffix`
- `job`, `catch_phrase`

### Internet
- `url`, `domain_name`
- `ipv4`, `ipv6`, `mac_address`
- `user_agent`, `slug`

### Dates & Times
- `date`, `date_time`
- `past_date`, `future_date`
- `time`, `unix_time`

### Text
- `text`, `sentence`, `paragraph`
- `word`, `words`

### Numbers
- `random_int`, `random_digit`
- `random_number`, `random_float`

### Special
- `person` - Complete person object
- `user` - User credentials object
- `random_from` - Pick from provided list
- `uuid` - Generate UUID

## 🗄️ Supported Databases

| Database   | Driver                 | Connection String Example                    |
|------------|------------------------|----------------------------------------------|
| SQLite     | `sqlite`               | `sqlite:///path/to/database.db`              |
| MySQL      | `mysql+mysqlconnector` | `mysql+mysqlconnector://user:pass@host/db`   |
| PostgreSQL | `postgresql+psycopg2`  | `postgresql+psycopg2://user:pass@host/db`    |

### SQLite Example
```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "myapp.db"
  }
}
```

### MySQL Example
```json
{
  "database": {
    "dbtype": "mysql",
    "drivername": "mysql+mysqlconnector",
    "username": "root",
    "password": "password",
    "host": "localhost",
    "port": 3306,
    "database": "myapp"
  }
}
```

### PostgreSQL Example
```json
{
  "database": {
    "dbtype": "postgresql",
    "drivername": "postgresql+psycopg2",
    "username": "postgres",
    "password": "password",
    "host": "localhost",
    "port": 5432,
    "database": "myapp"
  }
}
```

## 📚 Documentation

- **[Getting Started](docs/getting-started.md)** - Installation and basic usage
- **[Schema Reference](docs/schema-reference.md)** - Complete schema documentation
- **[Data Generators](docs/generators.md)** - All available data generators
- **[Database Support](docs/databases.md)** - Database-specific configuration
- **[API Reference](docs/api-reference.md)** - Python and Node.js APIs
- **[Examples](docs/examples.md)** - Real-world examples
- **[Troubleshooting](docs/troubleshooting.md)** - Common issues and solutions

### Language-Specific Docs
- **[Python Package](python/README.md)** - Python-specific documentation
- **[Node.js Package](node/README.md)** - Node.js/TypeScript documentation
- **[Go Core](golang/README.md)** - Go core documentation

## ⚡ Performance

Fakestack's Go core delivers exceptional performance:

| Dataset    | Python v2.0 | Go Core v2.1 | Speedup |
|------------|-------------|--------------|---------|
| 1K rows    | ~2.5s       | ~0.1s        | **25x** |
| 10K rows   | ~25s        | ~0.8s        | **31x** |
| 100K rows  | ~250s       | ~6s          | **42x** |

*Benchmarks run on: MacBook Pro M1, 16GB RAM, SQLite database*

## 📁 Repository Structure

```
fake-stack/
├── golang/           # Go core implementation
│   ├── *.go         # Source files
│   ├── Formula/     # Homebrew formula
│   └── README.md    # Go documentation
├── python/           # Python package (PyPI: fakestack)
│   ├── fakestack/   # Python module
│   ├── tests/       # Integration tests
│   └── README.md    # Python documentation
├── node/             # Node.js package (npm: fakestack)
│   ├── src/         # TypeScript source
│   ├── tests/       # Integration tests
│   └── README.md    # Node.js documentation
├── docs/             # Documentation
├── bin/              # Compiled binaries
└── scripts/          # Build scripts
```

## 🤝 Contributing

Contributions welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Development Setup

**Go Core:**
```bash
cd golang
go mod download
go build
go test -v ./...
```

**Python:**
```bash
cd python
pip install -e ".[dev]"
pytest tests/ -v
black fakestack/
```

**Node.js:**
```bash
cd node
npm install
npm test
npm run build
```

## 🧪 Testing

```bash
# Go tests
cd golang && go test -v ./...

# Python tests (all versions: 3.8-3.13)
cd python && pytest tests/ -v

# Node.js tests (Node 18+)
cd node && npm test
```

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

Built with:
- [gofakeit](https://github.com/brianvoe/gofakeit) - Go fake data generation
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver
- [go-mysql-driver](https://github.com/go-sql-driver/mysql) - MySQL driver
- [pq](https://github.com/lib/pq) - PostgreSQL driver

## 📞 Support

- 🐛 **Issues**: [GitHub Issues](https://github.com/0xdps/fake-stack/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/0xdps/fake-stack/discussions)
- 📦 **PyPI**: https://pypi.org/project/fakestack/
- 📦 **npm**: https://www.npmjs.com/package/fakestack
- 🍺 **Homebrew**: `brew install 0xdps/fakestack`

## 🔖 Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and release notes.

---

**Made with ❤️ by [Devendra Pratap](https://github.com/0xdps)**
