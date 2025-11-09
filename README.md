# Schemer - JSON Schema to Database Generator

**Schemer** is a Python tool that allows you to generate database tables and populate them with realistic fake data based on a provided JSON schema. It provides a simple command-line interface to quickly create and populate database tables without the need for manual schema definition.

## Features

- Convert JSON schema to database tables
- Populate tables with realistic fake data using Faker
- Supports MySQL, PostgreSQL, and SQLite databases
- Reference data between tables (foreign keys)
- Generate unique and consistent test data

## Installation

You can install Schemer using pip:

```bash
pip install git+https://github.com/0xdps/fake-db-generator.git
```

## Usage

After installing Schemer, you can use the `schemer` command followed by various options:

- `-c` / `--create-table`: Create the database tables
- `-p` / `--populate-data`: Populate data into the database
- `-f` / `--file <path>`: Specify the path to the JSON schema file
- `-d` / `--download-schema`: Download example schema
- `-h` / `--help`: Display help

### Example Usage

```bash
# Create tables from schema
schemer -c -f schema.json

# Populate tables with fake data
schemer -p -f schema.json

# Create and populate in one command
schemer -cpf schema.json

# Download example schema
schemer -d schema.json
```

### Supported Databases

Schemer supports:
- MySQL (via `mysql+mysqlconnector`)
- PostgreSQL (via `postgresql+psycopg2`)
- SQLite

## Example JSON Schemas

You can find example JSON schemas [here](schemer/data/).
