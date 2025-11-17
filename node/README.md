<p align="center">
    <img src="https://raw.githubusercontent.com/0xdps/fake-stack/trunk/assets/fake-stack.svg" alt="Fakestack Logo" width="200"/>
</p>

<h1 align="center">Fakestack</h1>

<p align="center">
High-performance database generator powered by Go. Generate database tables and populate them with realistic fake data from JSON schemas.
</p>

## Features

⚡ **10-50x faster** than pure JavaScript implementations  
🗄️ **Multi-database support** - SQLite, MySQL, PostgreSQL  
💡 **50+ data generators** - names, emails, addresses, and more  
🎯 **Simple API** - CLI and programmatic usage  
🌍 **Cross-platform** - Works on Linux, macOS, and Windows  
📦 **Zero dependencies** - Includes pre-compiled binaries

## Installation

```bash
npm install fakestack
# or
yarn add fakestack
# or
pnpm add fakestack
```

## 🎯 Quick Start

### CLI Usage

```bash
# Download example schema
npx fakestack -d .

# Create tables and populate with data
npx fakestack -c -p -f schema.json
```

### Programmatic Usage

```typescript
import { fakestack } from 'fakestack';

// Create tables and populate data
await fakestack({
  createTables: true,
  populateData: true,
  schemaFile: 'schema.json'
});

// Or download example schema
await fakestack({
  downloadSchema: '.'
});
```

### Advanced Usage

```typescript
import { runFakestack } from 'fakestack';

// Run with custom arguments
const exitCode = await runFakestack(['-c', '-p', '-f', 'schema.json']);

if (exitCode === 0) {
  console.log('Success!');
}
```

## 📖 CLI Options

```
-c, -create-table     Create database tables from schema
-p, -populate-data    Populate tables with fake data
-f, -file <path>      Path to JSON schema file
-d, -download-schema  Download example schema to specified path
```

## 📊 Example Schema

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
        {"name": "email", "generator": "email"}
      ]
    }
  ]
}
```

## 🎨 Available Generators

- **Personal**: name, first_name, last_name, email, username, password
- **Address**: address, city, country, street_address, postcode
- **Company**: company, job, catch_phrase
- **Internet**: url, domain_name, ipv4, ipv6, mac_address
- **Dates**: date, date_time, past_date, future_date
- **Text**: text, sentence, paragraph, word
- **Numbers**: random_int, random_digit, random_number
- And many more!

## 📊 Performance

| Rows | Node.js (faker) | fakestack | Speedup |
|------|----------------|-----------------|---------|
| 1,000 | ~3s | ~0.1s | **30x** ⚡ |
| 10,000 | ~30s | ~0.8s | **37x** ⚡ |
| 100,000 | ~300s | ~6s | **50x** ⚡ |

## 🗄️ Supported Databases

- SQLite (no additional setup required)
- MySQL / MariaDB
- PostgreSQL

## 🔧 TypeScript Support

Full TypeScript support with type definitions included:

```typescript
import { fakestack, FakestackOptions, runFakestack } from 'fakestack';

const options: FakestackOptions = {
  createTables: true,
  populateData: true,
  schemaFile: 'schema.json'
};

await fakestack(options);
```

## 📚 Documentation

📖 **[Full Documentation on ReadTheDocs](https://fake-stack.readthedocs.io/)** - Complete documentation with examples and tutorials

- [Main Repository](https://github.com/0xdps/fake-stack)
- [Python Package](https://pypi.org/project/fakestack/)
- [Contributing Guide](https://github.com/0xdps/fake-stack/blob/trunk/CONTRIBUTING.md)
- [Report Issues](https://github.com/0xdps/fake-stack/issues)

## Contributing

Contributions welcome! See [CONTRIBUTING.md](https://github.com/0xdps/fake-stack/blob/trunk/CONTRIBUTING.md)

## License

MIT - see [LICENSE](https://github.com/0xdps/fake-stack/blob/trunk/LICENSE)

## Built With

- [Go](https://golang.org/) - High-performance core
- [gofakeit](https://github.com/brianvoe/gofakeit) - Fake data generation
