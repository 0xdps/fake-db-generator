# Getting Started with Fakestack

This guide will help you get started with Fakestack in just a few minutes.

## Installation

Choose your preferred installation method:

### Python (pip)
```bash
pip install fakestack
```

### Node.js (npm)
```bash
npm install fakestack
```

### Homebrew (macOS/Linux)
```bash
brew tap 0xdps/packages
brew install fakestack
```

### Direct Binary Download
Download pre-built binaries from [GitHub Releases](https://github.com/0xdps/fake-stack/releases).

## Your First Database

### Step 1: Download Example Schema

```bash
fakestack -d .
```

This creates a `schema.json` file with an example database structure.

### Step 2: Generate the Database

```bash
fakestack -c -p -f schema.json
```

This will:
1. Create database tables based on the schema
2. Populate them with realistic fake data

### Step 3: Verify the Data

```bash
# For SQLite (default)
sqlite3 test.db "SELECT * FROM users LIMIT 5;"

# For MySQL
mysql -u root -p testdb -e "SELECT * FROM users LIMIT 5;"

# For PostgreSQL
psql -U postgres -d testdb -c "SELECT * FROM users LIMIT 5;"
```

## Using in Your Code

### Python

```python
from fakestack import fakestack

# Generate database
exit_code = fakestack(['-c', '-p', '-f', 'schema.json'])

if exit_code == 0:
    print("Database generated successfully!")
```

### Node.js / TypeScript

```javascript
import { fakestack } from 'fakestack';

// Generate database
const exitCode = await fakestack(['-c', '-p', '-f', 'schema.json']);

if (exitCode === 0) {
    console.log('Database generated successfully!');
}
```

## Next Steps

- **[Schema Reference](schema-reference.md)** - Learn how to create custom schemas
- **[Data Generators](generators.md)** - Explore all available data generators
- **[Database Support](databases.md)** - Configure different databases
- **[Examples](examples.md)** - See real-world examples

## Common Commands

```bash
# Download example schema
fakestack -d .

# Create tables only
fakestack -c -f schema.json

# Populate data only (tables must exist)
fakestack -p -f schema.json

# Create and populate in one command
fakestack -c -p -f schema.json

# Get help
fakestack -h
```

## Troubleshooting

### Binary Not Found
If you get a "binary not found" error, make sure you've installed fakestack correctly:
```bash
# Python
pip install --upgrade fakestack

# Node.js
npm install --save fakestack

# Homebrew
brew upgrade 0xdps/fakestack
```

### Permission Denied (Unix/macOS/Linux)
The binary might not have execute permissions:
```bash
chmod +x ~/.local/lib/python*/site-packages/fakestack/bin/fakestack-*
```

### Database Connection Errors
Verify your database configuration in the schema file:
- Check credentials (username, password)
- Verify host and port
- Ensure database exists
- Check network connectivity

For more help, see [Troubleshooting](troubleshooting.md).
