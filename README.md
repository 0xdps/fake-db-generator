# Fakestack - JSON Schema to Database Generator# Fakestack - JSON Schema to Database Generator# Fakestack - JSON Schema to Database Generator



[![PyPI version](https://badge.fury.io/py/fakestack.svg)](https://pypi.org/project/fakestack/)

[![Python Versions](https://img.shields.io/pypi/pyversions/fakestack.svg)](https://pypi.org/project/fakestack/)

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)[![PyPI version](https://badge.fury.io/py/fakestack.svg)](https://pypi.org/project/fakestack/)**Fakestack** is a Python tool that allows you to generate database tables and populate them with realistic fake data based on a provided JSON schema. It provides a simple command-line interface to quickly create and populate database tables without the need for manual schema definition.

[![CI Status](https://github.com/0xdps/fake-db-generator/workflows/Test%20%26%20Build/badge.svg)](https://github.com/0xdps/fake-db-generator/actions)

[![Python Versions](https://img.shields.io/pypi/pyversions/fakestack.svg)](https://pypi.org/project/fakestack/)

**Fakestack** is a Python tool that generates database tables and populates them with realistic fake data based on JSON schema definitions. Perfect for testing, development, and prototyping.

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)## Features

## Features

[![CI Status](https://github.com/0xdps/fake-db-generator/workflows/Test%20%26%20Build/badge.svg)](https://github.com/0xdps/fake-db-generator/actions)

- 🚀 **Schema-Driven** - Define tables and data rules in JSON

- 💡 **Realistic Data** - Uses Faker for authentic test data- Convert JSON schema to database tables

- 🗄️ **Multi-Database** - MySQL, PostgreSQL, and SQLite support

- 🔗 **Relationships** - Foreign key and referential data support**Fakestack** is a Python tool that allows you to generate database tables and populate them with realistic fake data based on a provided JSON schema. It provides a simple command-line interface to quickly create and populate database tables without the need for manual schema definition.- Populate tables with realistic fake data using Faker

- ⚡ **Fast Setup** - Create and populate databases in seconds

- Supports MySQL, PostgreSQL, and SQLite databases

## Installation

## ✨ Features- Reference data between tables (foreign keys)

```bash

pip install fakestack- Generate unique and consistent test data

```

- 🚀 **Schema-Driven**: Define tables and data generation rules in JSON

Or install from source:

- 💡 **Realistic Fake Data**: Uses Faker library for authentic-looking test data## Installation

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

## 🤝 Contributing

We welcome contributions! Please see:

- [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guidelines
- [DEVELOPMENT.md](DEVELOPMENT.md) - Development setup
- [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) - Community guidelines

### Quick Contribution Guide

```bash
# Fork and clone
git clone https://github.com/YOUR-USERNAME/fake-db-generator.git
cd fake-db-generator

# Create virtual environment
python -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate

# Install in development mode
pip install -e ".[dev]"

# Create feature branch
git checkout -b feature/amazing-feature

# Make changes and test
pytest
black .
isort .
flake8 .

# Commit and push
git commit -m "feat: add amazing feature"
git push origin feature/amazing-feature
```

## 🧪 Testing

Run tests with pytest:

```bash
# Run all tests
pytest

# With coverage
pytest --cov=fakestack --cov-report=html

# Run specific test file
pytest tests/test_schema.py
```

## 📄 License

This project is licensed under the MIT License - see [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Faker](https://github.com/joke2k/faker) for realistic data generation
- Uses [SQLAlchemy](https://www.sqlalchemy.org/) for database abstraction
- Uses [Pydantic](https://github.com/pydantic/pydantic) for data validation

## 📞 Support

- 🐛 **Issues**: [GitHub Issues](https://github.com/0xdps/fake-db-generator/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/0xdps/fake-db-generator/discussions)
- 📖 **Documentation**: [README](https://github.com/0xdps/fake-db-generator#readme)

## 🔖 Version History

See [CHANGELOG.md](CHANGELOG.md) for version history and release notes.

---

**Made with ❤️ by Devendra Pratap**
