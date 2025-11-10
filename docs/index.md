# Fakestack Documentation

![PyPI](https://img.shields.io/pypi/v/fakestack)
![npm](https://img.shields.io/npm/v/fakestack)
![Python Versions](https://img.shields.io/pypi/pyversions/fakestack)
![License](https://img.shields.io/github/license/0xdps/fake-stack)

**Fakestack** is a powerful tool for generating realistic fake data and populating databases with ease. Built with a high-performance Go core and accessible through Python, Node.js, and CLI interfaces.

## Features

✨ **Multi-Language Support** - Use from Python, Node.js, or command line  
🚀 **High Performance** - Go-powered core for blazing fast data generation  
🎲 **50+ Data Generators** - Names, emails, addresses, dates, and more  
💾 **Multi-Database** - SQLite, MySQL, and PostgreSQL support  
📊 **Realistic Data** - Generate contextually appropriate fake data  
🔗 **Foreign Keys** - Automatic relationship handling  
⚡ **Batch Operations** - Efficient bulk data insertion  
🎯 **Zero Dependencies** - Self-contained with bundled binaries  

## Quick Start

### Installation

=== "Python"
    ```bash
    pip install fakestack
    ```

=== "Node.js"
    ```bash
    npm install fakestack
    ```

=== "Homebrew"
    ```bash
    brew install 0xdps/fakestack/fakestack
    ```

### Basic Usage

=== "Python"
    ```python
    from fakestack import Fakestack

    schema = {
        "database": {
            "dbtype": "sqlite",
            "drivername": "sqlite",
            "database": "test.db"
        },
        "tables": [...],
        "populate": [...]
    }

    faker = Fakestack(schema)
    faker.run()
    ```

=== "Node.js"
    ```javascript
    const Fakestack = require('fakestack');

    const schema = {
        database: {
            dbtype: 'sqlite',
            drivername: 'sqlite',
            database: 'test.db'
        },
        tables: [...],
        populate: [...]
    };

    const faker = new Fakestack(schema);
    await faker.run();
    ```

=== "CLI"
    ```bash
    fakestack schema.json
    ```

## Documentation

- **[Getting Started](getting-started.md)** - Installation and first steps
- **[Schema Reference](schema-reference.md)** - Complete schema format
- **[Data Generators](generators.md)** - All available generators
- **[Database Support](databases.md)** - SQLite, MySQL, PostgreSQL guides
- **[API Reference](api-reference.md)** - Python and Node.js APIs
- **[Examples](examples.md)** - Real-world use cases
- **[Troubleshooting](troubleshooting.md)** - Common issues and solutions

## Supported Databases

| Database | Support | Notes |
|----------|---------|-------|
| SQLite | ✅ Full | File-based, no server required |
| MySQL | ✅ Full | Version 5.7+ |
| PostgreSQL | ✅ Full | Version 10+ |

## Platform Support

| Platform | Architecture | Status |
|----------|-------------|--------|
| macOS | Intel (x64) | ✅ Supported |
| macOS | Apple Silicon (ARM64) | ✅ Supported |
| Linux | x64 | ✅ Supported |
| Windows | x64 | ✅ Supported |
| Windows | x86 | ✅ Supported |

## Community

- **GitHub**: [0xdps/fake-stack](https://github.com/0xdps/fake-stack)
- **Issues**: [Report a bug](https://github.com/0xdps/fake-stack/issues)
- **Discussions**: [Ask questions](https://github.com/0xdps/fake-stack/discussions)

## License

Fakestack is released under the [MIT License](https://github.com/0xdps/fake-stack/blob/trunk/LICENSE).
