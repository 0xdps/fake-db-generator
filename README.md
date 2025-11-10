# Fakestack - High-Performance Database Generator

[![PyPI](https://img.shields.io/pypi/v/fakestack)](https://pypi.org/project/fakestack/)
[![npm](https://img.shields.io/npm/v/fakestack)](https://www.npmjs.com/package/fakestack)
[![Python](https://img.shields.io/pypi/pyversions/fakestack)](https://pypi.org/project/fakestack/)
[![Node](https://img.shields.io/node/v/fakestack)](https://www.npmjs.com/package/fakestack)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Generate databases from JSON schemas with realistic fake data. **10-100x faster** than pure Python/JavaScript implementations.

> 🚀 **v2.1.0**: Go core implementation with Python and Node.js wrappers - one package name across both ecosystems!

## ✨ Features

- 🚀 **Schema-Driven** - Define tables and data in JSON
- ⚡ **High Performance** - Go core for blazing speed
- 💡 **Realistic Data** - 50+ generators (names, emails, addresses, etc.)
- 🗄️ **Multi-Database** - SQLite, MySQL, PostgreSQL
- 🎯 **Simple API** - CLI and programmatic usage
- 🌍 **Cross-Platform** - Linux, macOS, Windows (amd64 & arm64)
- 📦 **Multi-Ecosystem** - Same package name for Python (pip) and Node.js (npm)
- 🔧 **Zero Dependencies** - Batteries included via bundled binaries

## 📦 Installation

**Python (PyPI):**
```bash
pip install fakestack
```

**Node.js (npm):**
```bash
npm install fakestack
```

**Homebrew (macOS/Linux):**
```bash
brew install 0xdps/fakestack
```

**Direct Binary Download:**
Download from [GitHub Releases](https://github.com/0xdps/fake-stack/releases)

## 🚀 Quick Start

### Python

```bash
# CLI usage
fakestack -d .                    # Download example schema
fakestack -c -p -f schema.json    # Create and populate

# Python API
from fakestack import fakestack
fakestack(['-c', '-p', '-f', 'schema.json'])
```

### Node.js / TypeScript

```bash
npm install fakestack

# CLI usage
npx fakestack -d .                # Download example schema  
npx fakestack -c -p -f schema.json # Create and populate

# Programmatic usage
const { fakestack } = require('fakestack');
await fakestack(['-c', '-p', '-f', 'schema.json']);
```

### Go (Core)

```bash
cd golang
go build -o fakestack
./fakestack -d .
./fakestack -c -p -f schema.json
```

## 📦 Installation

```bash

git clone https://github.com/0xdps/fake-db-generator.git- 🛡️ **Multi-Database Support**: Works with MySQL, PostgreSQL, and SQLite

cd fake-db-generator

pip install -e .- ⚡ **Relationship Support**: Generate data with foreign key relationshipsYou can install Fakestack using pip:

```

- 🎯 **Flexible Generators**: Custom data generators with parameter support

## Quick Start

- 🔄 **Repeatable**: Consistent data generation for testing```bash

Download example schema:

```bashpip install git+https://github.com/0xdps/fake-db-generator.git

fakestack -d

```## 📦 Installation```



Create tables and populate with data:

```bash

fakestack -cpf schema.json### From PyPI (Recommended)## Usage

```



Or separately:

```bash```bashAfter installing Fakestack, you can use the `fakestack` command followed by various options:

fakestack -c -f schema.json  # Create tables

fakestack -p -f schema.json  # Populate datapip install fakestack

```

```- `-c` / `--create-table`: Create the database tables

## Schema Example

- `-p` / `--populate-data`: Populate data into the database

```json

{### From Source- `-f` / `--file <path>`: Specify the path to the JSON schema file

  "database": {

    "dbtype": "sqlite",- `-d` / `--download-schema`: Download example schema

    "drivername": "sqlite",

    "database": "test.db"```bash- `-h` / `--help`: Display help

  },

  "tables": [git clone https://github.com/0xdps/fake-db-generator.git

    {

      "name": "users",cd fake-db-generator### Example Usage

      "columns": [

        {pip install -e .

          "name": "id",

          "type": "integer",``````bash

          "options": {"primary_key": true, "autoincrement": true}

        },# Create tables from schema

        {

          "name": "username",### Development Installationfakestack -c -f schema.json

          "type": {"name": "string", "args": {"length": 50}},

          "options": {"nullable": false, "unique": true}

        },

        {```bash# Populate tables with fake data

          "name": "email",

          "type": {"name": "string", "args": {"length": 100}},git clone https://github.com/0xdps/fake-db-generator.gitfakestack -p -f schema.json

          "options": {"nullable": false, "unique": true}

        }cd fake-db-generator

      ]

    }pip install -e ".[dev]"# Create and populate in one command

  ],

  "populate": [```fakestack -cpf schema.json

    {

      "name": "users",

      "count": 50,

      "fields": [## 🚀 Quick Start# Download example schema

        {"name": "username", "generator": "user_name"},

        {"name": "email", "generator": "email"}fakestack -d schema.json

      ]

    }### 1. Download Example Schema```

  ]

}

```

```bash### Supported Databases

## Supported Databases

fakestack -d

| Database   | Driver                 | Connection String                          |

|------------|------------------------|--------------------------------------------|```Fakestack supports:

| MySQL      | mysql+mysqlconnector   | `mysql+mysqlconnector://user:pass@host/db` |

| PostgreSQL | postgresql+psycopg2    | `postgresql+psycopg2://user:pass@host/db`  |- MySQL (via `mysql+mysqlconnector`)

| SQLite     | sqlite                 | `sqlite:///path/to/database.db`            |

This creates a `schema.json` file in your current directory.- PostgreSQL (via `postgresql+psycopg2`)

## Data Generators

- SQLite

Common Faker generators available:

### 2. Create Tables

- **Personal**: `first_name`, `last_name`, `email`, `user_name`, `password`

- **Address**: `address`, `city`, `country`, `street_address`## Example JSON Schemas

- **Company**: `company`, `job`, `catch_phrase`

- **Internet**: `url`, `domain_name`, `ipv4`, `mac_address````bash

- **Dates**: `date`, `date_time`, `past_date`, `future_date`

- **Text**: `text`, `sentence`, `paragraph`, `word`fakestack -c -f schema.jsonYou can find example JSON schemas [here](fakestack/data/).

- **Numbers**: `random_int`, `random_digit`

- **Custom**: `person`, `user`, `random_from`, `unique_item````



See [examples](fakestack/data/) for more complex schemas.### 3. Populate with Data



## Command-Line Options```bash

fakestack -p -f schema.json

``````

fakestack [-h] [-c] [-p] [-f FILE] [-d]

### 4. Or Do Both at Once

Options:

  -c, --create-table        Create database tables```bash

  -p, --populate-data       Populate tables with fake datafakestack -cpf schema.json

  -f, --file FILE           Path to JSON schema file```

  -d, --download-schema     Download example schema

  -h, --help                Show help message## 📖 Usage

```

### Command-Line Options

## Documentation

- `-c` / `--create-table` - Create database tables from schema

- **[Contributing Guide](CONTRIBUTING.md)** - How to contribute- `-p` / `--populate-data` - Populate tables with fake data

- **[Development Guide](DEVELOPMENT.md)** - Developer setup- `-f` / `--file <path>` - Specify path to JSON schema file

- **[Changelog](CHANGELOG.md)** - Version history- `-d` / `--download-schema` - Download example schema to current directory

- **[Examples](fakestack/data/)** - Example schemas- `-h` / `--help` - Display help message



## Contributing### Schema Format



We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for:```json

- How to submit bug reports and feature requests{

- Development setup and workflow  "database": {

- Code style guidelines    "dbtype": "mysql",

- Pull request process    "drivername": "mysql+mysqlconnector",

- Community guidelines and code of conduct    "username": "root",

    "password": "password",

Quick contribution:    "host": "localhost",

```bash    "database": "testdb"

git clone https://github.com/YOUR-USERNAME/fake-db-generator.git  },

cd fake-db-generator  "tables": [

pip install -e ".[dev]"    {

pytest  # Run tests      "name": "users",

```      "columns": [

        {

## Security          "name": "id",

          "type": "integer",

To report security vulnerabilities, please email **dps.manit@gmail.com** with details. Do not open public issues for security concerns.          "options": {"primary_key": true, "autoincrement": true}

        },

## License        {

          "name": "username",

This project is licensed under the MIT License - see [LICENSE](LICENSE) file for details.          "type": {"name": "string", "args": {"length": 50}},

          "options": {"nullable": false, "unique": true}

## Support        }

      ],

- 🐛 [Report Issues](https://github.com/0xdps/fake-db-generator/issues)      "indexes": []

- 💬 [Discussions](https://github.com/0xdps/fake-db-generator/discussions)    }

- 📧 Email: dps.manit@gmail.com  ],

  "populate": [

---    {

      "name": "users",

**Made with ❤️ by Devendra Pratap**      "count": 100,

      "fields": [
        {"name": "username", "generator": "user_name"},
        {"name": "email", "generator": "email"}
      ]
    }
  ]
}
```

## 🗄️ Supported Databases

| Database   | Driver                    | Connection String                          |
|------------|---------------------------|--------------------------------------------|
| MySQL      | `mysql+mysqlconnector`    | `mysql+mysqlconnector://user:pass@host/db` |
| PostgreSQL | `postgresql+psycopg2`     | `postgresql+psycopg2://user:pass@host/db`  |
| SQLite     | `sqlite`                  | `sqlite:///path/to/database.db`            |

## 🎨 Data Generators

Fakestack uses Faker under the hood. Available generators include:

- **Personal**: `first_name`, `last_name`, `email`, `user_name`, `password`
- **Address**: `address`, `city`, `country`, `street_address`, `postcode`
- **Company**: `company`, `job`, `catch_phrase`
- **Internet**: `url`, `domain_name`, `ipv4`, `mac_address`
- **Dates**: `date`, `date_time`, `past_date`, `future_date`
- **Text**: `text`, `sentence`, `paragraph`, `word`
- **Numbers**: `random_int`, `random_digit`, `random_number`
- **Custom**: `person` (returns Person object), `user` (returns User object)
- **Special**: `random_from` (pick from list), `unique_item` (unique values)

### Advanced Features

- **Foreign Keys**: Reference data from other tables
- **Unique Values**: Ensure no duplicates
- **Custom Objects**: Access nested properties (e.g., `person.first_name`)

## 📚 Documentation

- **[Getting Started](tutorials/01_getting_started.md)** - Installation and first steps
- **[Schema Guide](tutorials/02_schema_guide.md)** - Complete schema reference
- **[Examples](fakestack/data/)** - Example JSON schemas
- **[Contributing](CONTRIBUTING.md)** - How to contribute
- **[Development](DEVELOPMENT.md)** - Developer guide

## 💡 Examples

### Simple User Table

```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "users.db"
  },
  "tables": [
    {
      "name": "users",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true}},
        {"name": "name", "type": {"name": "string", "args": {"length": 100}}, "options": {}},
        {"name": "email", "type": {"name": "string", "args": {"length": 100}}, "options": {"unique": true}}
      ]
    }
  ],
  "populate": [
    {
      "name": "users",
      "count": 50,
      "fields": [
        {"name": "name", "generator": "name"},
        {"name": "email", "generator": "email"}
      ]
    }
  ]
}
```

More examples available in the [`fakestack/data/`](fakestack/data/) directory.

## 📁 Repository Structure

```
fake-stack/
├── golang/          # Go core implementation
│   ├── *.go        # Source files
│   ├── go.mod      # Dependencies
│   └── README.md   # Go-specific documentation
├── python/          # Python wrapper (PyPI: fakestack)
│   ├── fakestack/  # Python module
│   ├── tests/      # Integration tests
│   ├── pyproject.toml
│   └── README.md   # Python-specific documentation
├── node/            # Node.js wrapper (npm: fakestack)
│   ├── src/        # TypeScript source
│   ├── dist/       # Compiled JavaScript
│   ├── tests/      # Integration tests
│   ├── package.json
│   └── README.md   # Node-specific documentation
├── bin/             # Compiled binaries for all platforms
├── scripts/         # Build and release scripts
└── .github/         # CI/CD workflows
```

## 🤝 Contributing

Contributions welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Development Setup

**Go Core**:
```bash
cd golang
go mod download
go build
go test -v ./...
```

**Python**:
```bash
cd python
pip install -e ".[dev]"
pytest tests/ -v
black fakestack/
```

**Node.js**:
```bash
cd node
npm install
npm test
npm run build
```

## 🧪 Testing

Each language has its own test suite:

```bash
# Go tests
cd golang && go test -v ./...

# Python tests
cd python && pytest tests/ -v

# Node.js tests
cd node && npm test
```

CI/CD runs tests for all platforms on every push.

## 📄 License

MIT License - see [LICENSE](LICENSE) file

## 🙏 Built With

- [gofakeit](https://github.com/brianvoe/gofakeit) - Go fake data generation
- Go database drivers: go-sqlite3, go-mysql-driver, lib/pq

## 📞 Support

- 🐛 **Issues**: [GitHub Issues](https://github.com/0xdps/fake-stack/issues)
- 📦 **PyPI**: https://pypi.org/project/fakestack/
- 📦 **npm**: https://www.npmjs.com/package/fakestack
- 📖 **Documentation**: [README](https://github.com/0xdps/fake-db-generator#readme)

## 🔖 Version History

See [CHANGELOG.md](CHANGELOG.md) for version history and release notes.

---

**Made with ❤️ by Devendra Pratap**
