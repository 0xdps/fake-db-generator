# API Reference

Complete API documentation for Python and Node.js packages.

## Python API

### Installation

```bash
pip install fakestack
```

### Module Import

```python
from fakestack import Fakestack
```

### Fakestack Class

#### `Fakestack(schema: dict)`

Create a new Fakestack instance.

**Parameters:**
- `schema` (dict): JSON schema defining database, tables, and population

**Returns:**
- `Fakestack`: Instance ready to generate data

**Example:**
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
```

#### `run() -> None`

Execute the schema to create tables and populate data.

**Returns:**
- `None`

**Raises:**
- `RuntimeError`: If binary execution fails
- `FileNotFoundError`: If binary not found
- `ValueError`: If schema is invalid

**Example:**
```python
faker = Fakestack(schema)
faker.run()
```

#### `from_file(filepath: str) -> Fakestack`

Load schema from a JSON file.

**Parameters:**
- `filepath` (str): Path to JSON schema file

**Returns:**
- `Fakestack`: Instance with loaded schema

**Raises:**
- `FileNotFoundError`: If file doesn't exist
- `json.JSONDecodeError`: If JSON is invalid

**Example:**
```python
faker = Fakestack.from_file('schema.json')
faker.run()
```

### Complete Python Example

```python
import json
from fakestack import Fakestack

# Method 1: Direct schema
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
                {
                    "name": "id",
                    "type": "integer",
                    "options": {"primary_key": True, "autoincrement": True}
                },
                {
                    "name": "username",
                    "type": {"name": "string", "args": {"length": 50}},
                    "options": {"nullable": False, "unique": True}
                },
                {
                    "name": "email",
                    "type": {"name": "string", "args": {"length": 100}},
                    "options": {"nullable": False, "unique": True}
                }
            ],
            "indexes": []
        }
    ],
    "populate": [
        {
            "name": "users",
            "count": 100,
            "fields": [
                {"name": "username", "generator": "user_name"},
                {"name": "email", "generator": "email"}
            ]
        }
    ]
}

faker = Fakestack(schema)
faker.run()
print("✓ Database created with 100 users")

# Method 2: From file
faker = Fakestack.from_file('schema.json')
faker.run()
print("✓ Schema loaded and executed")
```

### Error Handling

```python
from fakestack import Fakestack

try:
    faker = Fakestack(schema)
    faker.run()
    print("✓ Success")
except RuntimeError as e:
    print(f"✗ Execution failed: {e}")
except FileNotFoundError as e:
    print(f"✗ Binary not found: {e}")
except ValueError as e:
    print(f"✗ Invalid schema: {e}")
```

### Type Hints

```python
from typing import Dict, Any
from fakestack import Fakestack

def generate_test_data(schema: Dict[str, Any]) -> None:
    """Generate test data from schema."""
    faker: Fakestack = Fakestack(schema)
    faker.run()

def load_and_generate(filepath: str) -> None:
    """Load schema from file and generate data."""
    faker: Fakestack = Fakestack.from_file(filepath)
    faker.run()
```

---

## Node.js API

### Installation

```bash
npm install fakestack
```

### Module Import

```javascript
// CommonJS
const Fakestack = require('fakestack');

// ES Modules (if configured)
import Fakestack from 'fakestack';
```

### Fakestack Class

#### `new Fakestack(schema)`

Create a new Fakestack instance.

**Parameters:**
- `schema` (Object): JSON schema defining database, tables, and population

**Returns:**
- `Fakestack`: Instance ready to generate data

**Example:**
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
```

#### `run()`

Execute the schema to create tables and populate data.

**Returns:**
- `Promise<void>`: Resolves when complete

**Throws:**
- `Error`: If binary execution fails
- `Error`: If binary not found
- `Error`: If schema is invalid

**Example:**
```javascript
const faker = new Fakestack(schema);
await faker.run();
```

#### `Fakestack.fromFile(filepath)`

Load schema from a JSON file.

**Parameters:**
- `filepath` (string): Path to JSON schema file

**Returns:**
- `Fakestack`: Instance with loaded schema

**Throws:**
- `Error`: If file doesn't exist or JSON is invalid

**Example:**
```javascript
const faker = Fakestack.fromFile('schema.json');
await faker.run();
```

### Complete Node.js Example

```javascript
const Fakestack = require('fakestack');
const fs = require('fs');

// Method 1: Direct schema
async function createDatabase() {
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
                    {
                        name: 'id',
                        type: 'integer',
                        options: { primary_key: true, autoincrement: true }
                    },
                    {
                        name: 'username',
                        type: { name: 'string', args: { length: 50 } },
                        options: { nullable: false, unique: true }
                    },
                    {
                        name: 'email',
                        type: { name: 'string', args: { length: 100 } },
                        options: { nullable: false, unique: true }
                    }
                ],
                indexes: []
            }
        ],
        populate: [
            {
                name: 'users',
                count: 100,
                fields: [
                    { name: 'username', generator: 'user_name' },
                    { name: 'email', generator: 'email' }
                ]
            }
        ]
    };

    const faker = new Fakestack(schema);
    await faker.run();
    console.log('✓ Database created with 100 users');
}

// Method 2: From file
async function loadAndGenerate() {
    const faker = Fakestack.fromFile('schema.json');
    await faker.run();
    console.log('✓ Schema loaded and executed');
}

// Execute
createDatabase().catch(console.error);
```

### Error Handling

```javascript
const Fakestack = require('fakestack');

async function generateData(schema) {
    try {
        const faker = new Fakestack(schema);
        await faker.run();
        console.log('✓ Success');
    } catch (error) {
        console.error('✗ Error:', error.message);
        throw error;
    }
}

// With promise chain
const faker = new Fakestack(schema);
faker.run()
    .then(() => console.log('✓ Success'))
    .catch(err => console.error('✗ Error:', err));
```

### TypeScript Support

```typescript
import Fakestack from 'fakestack';

interface Schema {
    database: {
        dbtype: string;
        drivername: string;
        database: string;
    };
    tables: any[];
    populate: any[];
}

async function generateTestData(schema: Schema): Promise<void> {
    const faker = new Fakestack(schema);
    await faker.run();
}

async function loadAndGenerate(filepath: string): Promise<void> {
    const faker = Fakestack.fromFile(filepath);
    await faker.run();
}
```

---

## CLI Usage

### Python CLI

After installing via pip, use the `fakestack` command:

```bash
# Basic usage
fakestack schema.json

# Using python -m
python -m fakestack schema.json
```

### Node.js CLI

After installing via npm, use the `fakestack` command:

```bash
# Global installation
npm install -g fakestack
fakestack schema.json

# Local installation
npx fakestack schema.json

# Or via package.json script
{
  "scripts": {
    "generate": "fakestack schema.json"
  }
}
npm run generate
```

### Homebrew CLI

After installing via Homebrew:

```bash
# Install
brew install 0xdps/fakestack

# Use
fakestack schema.json
```

---

## Common Patterns

### Environment-Based Configuration

**Python:**
```python
import os
from fakestack import Fakestack

env = os.getenv('ENV', 'development')

schema = {
    "database": {
        "dbtype": "sqlite",
        "drivername": "sqlite",
        "database": f"{env}.db"
    },
    # ...
}

faker = Fakestack(schema)
faker.run()
```

**Node.js:**
```javascript
const Fakestack = require('fakestack');

const env = process.env.ENV || 'development';

const schema = {
    database: {
        dbtype: 'sqlite',
        drivername: 'sqlite',
        database: `${env}.db`
    },
    // ...
};

const faker = new Fakestack(schema);
await faker.run();
```

### Multiple Schemas

**Python:**
```python
from fakestack import Fakestack

schemas = ['users.json', 'products.json', 'orders.json']

for schema_file in schemas:
    faker = Fakestack.from_file(schema_file)
    faker.run()
    print(f"✓ Generated {schema_file}")
```

**Node.js:**
```javascript
const Fakestack = require('fakestack');

const schemas = ['users.json', 'products.json', 'orders.json'];

for (const schemaFile of schemas) {
    const faker = Fakestack.fromFile(schemaFile);
    await faker.run();
    console.log(`✓ Generated ${schemaFile}`);
}
```

### Custom Record Counts

**Python:**
```python
import json
from fakestack import Fakestack

with open('schema.json', 'r') as f:
    schema = json.load(f)

# Adjust counts dynamically
for table in schema['populate']:
    table['count'] = 1000  # Generate 1000 records per table

faker = Fakestack(schema)
faker.run()
```

**Node.js:**
```javascript
const Fakestack = require('fakestack');
const fs = require('fs');

const schema = JSON.parse(fs.readFileSync('schema.json', 'utf8'));

// Adjust counts dynamically
schema.populate.forEach(table => {
    table.count = 1000;  // Generate 1000 records per table
});

const faker = new Fakestack(schema);
await faker.run();
```

---

## Performance Considerations

1. **Batch Size**: Fakestack uses optimized batch inserts (default: 1000 records per batch)
2. **Memory Usage**: Memory usage scales with record count and complexity
3. **Disk I/O**: SQLite performance depends on disk speed
4. **Network Latency**: MySQL/PostgreSQL performance depends on network speed
5. **Connection Pooling**: Reuse Fakestack instances when possible

### Benchmarks

| Records | SQLite | MySQL | PostgreSQL |
|---------|--------|-------|------------|
| 1,000   | ~0.5s  | ~1.0s | ~1.2s      |
| 10,000  | ~3.5s  | ~8.0s | ~9.5s      |
| 100,000 | ~32s   | ~75s  | ~85s       |
| 1,000,000 | ~5m  | ~12m  | ~14m       |

*Benchmarks on MacBook Pro M1, 16GB RAM*

---

## See Also

- [Schema Reference](schema-reference.md) - Complete schema format
- [Data Generators](generators.md) - All available generators
- [Examples](examples.md) - Real-world usage examples
- [Troubleshooting](troubleshooting.md) - Common issues and solutions
