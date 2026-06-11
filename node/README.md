<p align="center">
    <img src="https://raw.githubusercontent.com/0xdps/fake-stack/trunk/assets/fake-stack.svg" alt="Fakestack Logo" width="200"/>
</p>

<h1 align="center">Fakestack</h1>

<p align="center">
High-performance database generator powered by Go. Generate database tables and populate them with realistic fake data from JSON schemas.
</p>

## Features

⚡ **10-50x faster** than pure JavaScript implementations  
🗄️ **Multi-database support** - SQLite, MySQL, PostgreSQL, MariaDB, MS SQL Server, CockroachDB  
💡 **116+ data generators** - financial, products, animals, food, vehicles, books, and more  
🎨 **Template generator** - Custom data patterns with modifiers  
🎯 **Interactive generator** - Built-in schema generator with 10 pre-built templates  
🎯 **Simple API** - CLI and programmatic usage  
🌍 **Cross-platform** - Works on Linux, macOS, and Windows  
📦 **Tiny package** - <100KB (downloads binary on first run)  
🔄 **Auto-updates** - Checks for latest version on each run

## Installation

```bash
npm install fakestack
# or
yarn add fakestack
# or
pnpm add fakestack
```

> **Note**: The package automatically downloads the appropriate binary for your platform on first run (~10-20MB) and checks for updates on subsequent runs. The binary is cached in `~/.fakestack/bin/`.

## 🎯 Quick Start

### CLI Usage

```bash
# Generate schema interactively
npx fakestack -g .

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
-g, -generate <file>  Generate schema interactively (use '.' for default filename)
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

## 🎨 Available Generators (116+)

- **Personal & Identifiers**: name, first_name, last_name, email, username, password, ssn, ein
- **Financial & Payment**: credit_card, credit_card_type, credit_card_cvv, credit_card_exp, currency, bitcoin_address, iban, routing_number, price
- **Address & Location**: address, city, country, country_code, street_address, postcode, latitude, longitude, timezone, language, locale
- **Company & Job**: company, job, catch_phrase
- **Internet & Technology**: url, domain_name, ipv4, ipv6, mac_address, user_agent, slug, app_name, app_version
- **Dates & Times**: date, date_time, timestamp, past_date, future_date, time, year, month, weekday
- **Text & Content**: text, sentence, paragraph, word, quote, phrase, question, emoji
- **Numbers & Ranges**: random_int, integer (with min/max), random_digit, random_number, float, decimal (with min/max)
- **Products & E-commerce**: product_name, product_category, product_description, color, hex_color, price
- **Files & Media**: filename, file_extension, mime_type, image_url
- **Books & Entertainment**: book_title, book_author, book_genre, movie_name, movie_genre
- **Animals & Nature**: animal, animal_type, pet_name, cat, dog, bird, farm_animal
- **Food & Drink**: fruit, vegetable, breakfast, lunch, dinner, snack, dessert, drink
- **Vehicles**: car_maker, car_model, car_type, car_fuel_type, car_transmission_type
- **Template Generator**: Create custom patterns with modifiers (upper, lower, title, trim, truncate)

## 📊 Performance

| Rows | Node.js (faker) | fakestack | Speedup |
|------|----------------|-----------------|---------|
| 1,000 | ~3s | ~0.1s | **30x** ⚡ |
| 10,000 | ~30s | ~0.8s | **37x** ⚡ |
| 100,000 | ~300s | ~6s | **50x** ⚡ |

## 🗄️ Supported Databases

- **SQLite** - No additional setup required
- **MySQL** - Industry-standard relational database
- **MariaDB** - MySQL-compatible with enhanced features
- **PostgreSQL** - Advanced open-source database
- **MS SQL Server** - Microsoft's enterprise database
- **CockroachDB** - Distributed SQL database (PostgreSQL-compatible)

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
