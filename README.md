<p align="center">
  <img src="assets/fake-stack.svg" alt="Fakestack Logo" width="200"/>
</p>

<h1 align="center">Fakestack</h1>

<p align="center">

[![PyPI](https://img.shields.io/pypi/v/fakestack)](https://pypi.org/project/fakestack/)
[![npm](https://img.shields.io/npm/v/fakestack)](https://www.npmjs.com/package/fakestack)
[![Python](https://img.shields.io/badge/python-%3E%3D3.8-blue)](https://pypi.org/project/fakestack/)
[![Node](https://img.shields.io/node/v/fakestack)](https://www.npmjs.com/package/fakestack)
[![Documentation](https://readthedocs.org/projects/fake-stack/badge/?version=latest)](https://fake-stack.readthedocs.io/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xdps/fake-stack)](https://goreportcard.com/report/github.com/0xdps/fake-stack)

</p>

<p align="center">
  <strong>High-performance database generator with realistic fake data</strong>
</p>

Generate databases from JSON schemas with realistic fake data. **10-50x faster** than pure Python/JavaScript implementations thanks to a Go core with zero-dependency wrappers for Python and Node.js.

## ✨ Features

- 🚀 **Schema-Driven** - Define tables and data in simple JSON format
- ⚡ **High Performance** - Go core delivers 10-50x speed improvement
- 💡 **Realistic Data** - 116+ generators covering financial, localization, products, animals, food, vehicles, books, and more
- 🎨 **Custom Patterns** - Template generator for custom data formats (SKUs, IDs, codes, license plates)
- 🎯 **Interactive Generator** - Built-in schema generator with 10 pre-built templates
- 🗄️ **Multi-Database** - Works with SQLite, MySQL, PostgreSQL, MariaDB, MS SQL Server, CockroachDB
- 🎯 **Simple API** - Easy CLI and programmatic usage
- 🌍 **Cross-Platform** - Linux, macOS, Windows (amd64 & arm64)
- 📦 **Multi-Ecosystem** - Available on PyPI, npm, and Homebrew
- 🔧 **Zero Dependencies** - Self-contained with bundled binaries

## 📦 Installation

Choose your preferred package manager:

**Python (pip)** - Requires Python >= 3.8
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

### 1. Generate a Schema (Interactive)

Use the built-in interactive generator:

```bash
fakestack -g .
# or specify output filename
fakestack -g my-schema.json
```

Choose from 10 pre-built templates:
- **Users** - Basic user management
- **Employees** - Employee records with departments
- **Products** - E-commerce products
- **Orders** - Order management system
- **Customers** - Customer database
- **Blog Posts** - Content management
- **Inventory** - Stock management
- **Transactions** - Financial records
- **Students** - Educational records
- **Tasks** - Task management

Or download a basic example:

```bash
fakestack -d .
```

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
  -g, --generate <file>  Generate schema interactively (use '.' for default filename)
  -d, --download-schema  Download example schema
  -v, --version          Show version information
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

## 🎨 Available Generators (116+)

### Personal Data
- `first_name`, `last_name`, `name`
- `email`, `user_name`, `password`
- `phone_number`, `ssn`, `ein`

### Address & Location
- `address`, `street_address`
- `city`, `state`, `country`, `country_code`
- `postcode`, `latitude`, `longitude`
- `timezone`, `language`, `locale`

### Financial & Payment
- `credit_card`, `credit_card_type`, `credit_card_cvv`, `credit_card_exp`
- `currency`, `currency_long`, `price`
- `bitcoin_address`, `bitcoin_private_key`
- `iban`, `routing_number`

### Company & Job
- `company`, `company_suffix`
- `job`, `catch_phrase`

### Internet & Technology
- `url`, `domain_name`, `slug`
- `ipv4`, `ipv6`, `mac_address`
- `user_agent`, `chrome_user_agent`, `firefox_user_agent`, `safari_user_agent`, `opera_user_agent`
- `app_name`, `app_version`, `app_author`

### Dates & Times
- `date`, `date_time`, `timestamp`
- `past_date`, `future_date`
- `time`, `unix_time`
- `year`, `month`, `month_string`, `weekday`

### Text & Content
- `text`, `sentence`, `paragraph`
- `word`, `words`
- `quote`, `phrase`, `question`
- `emoji`, `emoji_description`, `emoji_category`

### Numbers & Ranges
- `random_int`, `integer` (with min/max range)
- `random_digit`, `random_number`
- `random_float`, `float`, `decimal` (with min/max range)

### Products & E-commerce
- `product_name`, `product_category`, `product_description`, `product_feature`
- `color`, `hex_color`, `safe_color`
- `price`

### Files & Media
- `filename`, `file_extension`, `mime_type`
- `image_url`

### Books & Entertainment
- `book_title`, `book_author`, `book_genre`
- `movie_name`, `movie_genre`

### Animals & Nature
- `animal`, `animal_type`, `pet_name`
- `cat`, `dog`, `bird`, `farm_animal`

### Food & Drink
- `fruit`, `vegetable`
- `breakfast`, `lunch`, `dinner`, `snack`, `dessert`
- `drink`

### Vehicles & Transportation
- `car_maker`, `car_model`, `car_type`
- `car_fuel_type`, `car_transmission_type`

### Special Generators
- `person` - Complete person object
- `user` - User credentials object
- `random_from` - Pick from provided list
- `uuid` - Generate UUID
- `template` - Custom patterns with modifiers

## 🗄️ Supported Databases

| Database      | Driver                 | Connection String Example                    |
|---------------|------------------------|----------------------------------------------|
| SQLite        | `sqlite`               | `sqlite:///path/to/database.db`              |
| MySQL         | `mysql+mysqlconnector` | `mysql+mysqlconnector://user:pass@host/db`   |
| PostgreSQL    | `postgresql+psycopg2`  | `postgresql+psycopg2://user:pass@host/db`    |
| MariaDB       | `mysql+mysqlconnector` | `mysql+mysqlconnector://user:pass@host/db`   |
| MS SQL Server | `mssql+pyodbc`         | `mssql+pyodbc://user:pass@host/db`           |
| CockroachDB   | `postgresql+psycopg2`  | `postgresql+psycopg2://user:pass@host:26257/db` |

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

## 🎨 Custom Data Patterns

Create custom data formats using the **template generator**:

```json
{
  "name": "sku",
  "generator": "template",
  "args": {
    "pattern": "{{word|upper|truncate(3)}}-{{random_int(1000,9999)}}"
  }
}
```
**Output**: `PRD-4821`, `INV-9234`, `STO-1456`

**More Examples:**
```json
{"pattern": "EMP-{{random_int(10000,99999)}}"}         // EMP-45678
{"pattern": "{{uuid|upper|truncate(8)}}"}               // A1B2C3D4
{"pattern": "{{city}}, {{state}} {{postcode}}"}         // San Francisco, CA 94102
{"pattern": "{{car_maker}}-{{car_model|upper}}"}        // Toyota-CAMRY
{"pattern": "{{country_code}}-{{random_int(100,999)}}"} // US-432
```

**Supported Modifiers**: `upper`, `lower`, `title`, `trim`, `truncate(n)`

📚 **[Template Examples](docs/TEMPLATE_EXAMPLES.md)** - Comprehensive examples  
📖 **[Custom Generators Guide](docs/CUSTOM_GENERATORS.md)** - Full guide with advanced patterns

## 📚 Documentation

📖 **[Full Documentation on ReadTheDocs](https://fake-stack.readthedocs.io/)** - Complete documentation with examples and tutorials

- **[Getting Started](docs/getting-started.md)** - Installation and basic usage
- **[Schema Reference](docs/schema-reference.md)** - Complete schema documentation
- **[Data Generators](docs/generators.md)** - All available data generators
- **[Custom Generators](docs/CUSTOM_GENERATORS.md)** - Template generator guide
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

- � **Documentation**: [ReadTheDocs](https://fake-stack.readthedocs.io/)
- �🐛 **Issues**: [GitHub Issues](https://github.com/0xdps/fake-stack/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/0xdps/fake-stack/discussions)
- 📦 **PyPI**: https://pypi.org/project/fakestack/
- 📦 **npm**: https://www.npmjs.com/package/fakestack
- 🍺 **Homebrew**: `brew tap 0xdps/packages && brew install fakestack`

## 🔖 Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and release notes.

---

**Made with ❤️ by [Devendra Pratap](https://github.com/0xdps)**
