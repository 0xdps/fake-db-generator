---
hide:
  - navigation
  - toc
---

<div class="hero" markdown>

<p align="center">
    <img src="https://raw.githubusercontent.com/0xdps/fake-stack/trunk/assets/fake-stack.svg" alt="Fakestack Logo" width="250"/>
</p>

# Fakestack

### Generate realistic fake data and populate databases with ease

<div class="badges" markdown>
[![PyPI](https://img.shields.io/pypi/v/fakestack?style=flat-square)](https://pypi.org/project/fakestack/)
[![npm](https://img.shields.io/npm/v/fakestack?style=flat-square)](https://www.npmjs.com/package/fakestack)
[![Python](https://img.shields.io/pypi/pyversions/fakestack?style=flat-square)](https://pypi.org/project/fakestack/)
[![License](https://img.shields.io/github/license/0xdps/fake-stack?style=flat-square)](https://github.com/0xdps/fake-stack/blob/trunk/LICENSE)
</div>

<div class="grid cards" markdown>

-   :material-language-python:{ .lg .middle } __Python__

    ---

    ```bash
    pip install fakestack
    ```

    [:octicons-arrow-right-24: Get started](getting-started.md)

-   :material-language-javascript:{ .lg .middle } __Node.js__

    ---

    ```bash
    npm install fakestack
    ```

    [:octicons-arrow-right-24: Get started](getting-started.md)

-   :material-package-variant:{ .lg .middle } __Homebrew__

    ---

    ```bash
    brew install 0xdps/fakestack
    ```

    [:octicons-arrow-right-24: Get started](getting-started.md)

-   :material-console:{ .lg .middle } __CLI__

    ---

    ```bash
    fakestack schema.json
    ```

    [:octicons-arrow-right-24: View examples](examples.md)

</div>

</div>

## Why Fakestack?

<div class="grid cards" markdown>

-   :material-rocket-launch:{ .lg .middle } __High Performance__

    ---

    Go-powered core for blazing fast data generation. Generate millions of records in seconds.

-   :material-database:{ .lg .middle } __Multi-Database__

    ---

    Full support for SQLite, MySQL, and PostgreSQL with automatic schema creation.

-   :material-puzzle:{ .lg .middle } __50+ Generators__

    ---

    Comprehensive data generators for names, emails, addresses, dates, and much more.

-   :material-link-variant:{ .lg .middle } __Smart Relations__

    ---

    Automatic foreign key handling with realistic data relationships.

-   :material-language-python:{ .lg .middle } __Multi-Language__

    ---

    Use from Python, Node.js, CLI, or directly as a Go library.

-   :material-package-down:{ .lg .middle } __Zero Dependencies__

    ---

    Self-contained with pre-built binaries. No external dependencies required.

</div>

## Quick Example

=== "Python"

    ```python
    from fakestack import Fakestack

    schema = {
        "database": {
            "dbtype": "sqlite",
            "drivername": "sqlite",
            "database": "users.db"
        },
        "tables": [
            {
                "name": "users",
                "columns": [
                    {"name": "id", "type": "integer", 
                     "options": {"primary_key": true, "autoincrement": true}},
                    {"name": "name", "type": {"name": "string", "args": {"length": 100}}, 
                     "options": {"nullable": false}},
                    {"name": "email", "type": {"name": "string", "args": {"length": 100}}, 
                     "options": {"nullable": false, "unique": true}}
                ],
                "indexes": []
            }
        ],
        "populate": [
            {
                "name": "users",
                "count": 1000,
                "fields": [
                    {"name": "name", "generator": "name"},
                    {"name": "email", "generator": "email"}
                ]
            }
        ]
    }

    faker = Fakestack(schema)
    faker.run()
    # ✓ Database created with 1000 users!
    ```

=== "Node.js"

    ```javascript
    const Fakestack = require('fakestack');

    const schema = {
        database: {
            dbtype: 'sqlite',
            drivername: 'sqlite',
            database: 'users.db'
        },
        tables: [
            {
                name: 'users',
                columns: [
                    {name: 'id', type: 'integer', 
                     options: {primary_key: true, autoincrement: true}},
                    {name: 'name', type: {name: 'string', args: {length: 100}}, 
                     options: {nullable: false}},
                    {name: 'email', type: {name: 'string', args: {length: 100}}, 
                     options: {nullable: false, unique: true}}
                ],
                indexes: []
            }
        ],
        populate: [
            {
                name: 'users',
                count: 1000,
                fields: [
                    {name: 'name', generator: 'name'},
                    {name: 'email', generator: 'email'}
                ]
            }
        ]
    };

    const faker = new Fakestack(schema);
    await faker.run();
    // ✓ Database created with 1000 users!
    ```

=== "CLI"

    ```bash
    # Create schema.json with your configuration
    fakestack schema.json
    # ✓ Database created with 1000 users!
    ```

## Explore Documentation

<div class="grid cards" markdown>

-   :material-book-open-page-variant:{ .lg .middle } __Getting Started__

    ---

    Learn the basics and create your first database in minutes.

    [:octicons-arrow-right-24: Start tutorial](getting-started.md)

-   :material-file-document:{ .lg .middle } __Schema Reference__

    ---

    Complete guide to JSON schema format and configuration.

    [:octicons-arrow-right-24: View reference](schema-reference.md)

-   :material-creation:{ .lg .middle } __Data Generators__

    ---

    Browse all 50+ available data generators with examples.

    [:octicons-arrow-right-24: See generators](generators.md)

-   :material-database-cog:{ .lg .middle } __Database Support__

    ---

    Setup guides for SQLite, MySQL, and PostgreSQL.

    [:octicons-arrow-right-24: Database guides](databases.md)

-   :material-code-braces:{ .lg .middle } __API Reference__

    ---

    Complete Python and Node.js API documentation.

    [:octicons-arrow-right-24: API docs](api-reference.md)

-   :material-file-code:{ .lg .middle } __Examples__

    ---

    Real-world schemas for e-commerce, blogs, and more.

    [:octicons-arrow-right-24: View examples](examples.md)

</div>

---

<div class="center-text" markdown>

### Ready to generate some data?

[Get Started](getting-started.md){ .md-button .md-button--primary .md-button--lg }
[View on GitHub](https://github.com/0xdps/fake-stack){ .md-button .md-button--lg }

</div>
