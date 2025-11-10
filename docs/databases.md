# Database Support

Fakestack supports three major database systems with native drivers and optimized connections.

## SQLite

### Overview

SQLite is a lightweight, file-based database perfect for development, testing, and small applications.

**Pros:**
- No server setup required
- Single file storage
- Fast for small to medium datasets
- Cross-platform compatible

**Cons:**
- Limited concurrent write access
- No network access
- Not suitable for high-concurrency applications

### Configuration

```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "path/to/database.db"
  }
}
```

### Connection Options

| Option | Required | Default | Description |
|--------|----------|---------|-------------|
| `dbtype` | Yes | - | Must be `"sqlite"` |
| `drivername` | Yes | - | Must be `"sqlite"` |
| `database` | Yes | - | Path to database file |

### File Paths

**Relative path:**
```json
"database": "myapp.db"
```

**Absolute path:**
```json
"database": "/Users/username/databases/myapp.db"
```

**In-memory (temporary):**
```json
"database": ":memory:"
```

### Usage Example

**Python:**
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

**Node.js:**
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
faker.run();
```

**CLI:**
```bash
fakestack schema.json
```

## MySQL

### Overview

MySQL is a popular open-source relational database management system ideal for web applications.

**Pros:**
- High performance
- Mature and widely supported
- Good for high-concurrency
- Scales well

**Cons:**
- Requires server setup
- More resource intensive
- More complex configuration

### Configuration

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

### Connection Options

| Option | Required | Default | Description |
|--------|----------|---------|-------------|
| `dbtype` | Yes | - | Must be `"mysql"` |
| `drivername` | Yes | - | Must be `"mysql+mysqlconnector"` |
| `username` | Yes | - | MySQL username |
| `password` | Yes | - | MySQL password |
| `host` | Yes | `"localhost"` | Server host |
| `port` | No | `3306` | Server port |
| `database` | Yes | - | Database name |

### Prerequisites

1. **Install MySQL Server:**
   ```bash
   # macOS (Homebrew)
   brew install mysql
   brew services start mysql
   
   # Ubuntu/Debian
   sudo apt-get install mysql-server
   sudo systemctl start mysql
   
   # Windows
   # Download from: https://dev.mysql.com/downloads/installer/
   ```

2. **Create Database:**
   ```sql
   mysql -u root -p
   CREATE DATABASE myapp;
   ```

3. **Create User (optional):**
   ```sql
   CREATE USER 'myuser'@'localhost' IDENTIFIED BY 'mypassword';
   GRANT ALL PRIVILEGES ON myapp.* TO 'myuser'@'localhost';
   FLUSH PRIVILEGES;
   ```

### Connection Examples

**Local development:**
```json
{
  "database": {
    "dbtype": "mysql",
    "drivername": "mysql+mysqlconnector",
    "username": "root",
    "password": "password",
    "host": "localhost",
    "port": 3306,
    "database": "myapp_dev"
  }
}
```

**Remote server:**
```json
{
  "database": {
    "dbtype": "mysql",
    "drivername": "mysql+mysqlconnector",
    "username": "app_user",
    "password": "secure_password",
    "host": "db.example.com",
    "port": 3306,
    "database": "production_db"
  }
}
```

**Docker container:**
```json
{
  "database": {
    "dbtype": "mysql",
    "drivername": "mysql+mysqlconnector",
    "username": "root",
    "password": "root",
    "host": "127.0.0.1",
    "port": 3306,
    "database": "test_db"
  }
}
```

### Docker Setup

```bash
# Start MySQL container
docker run -d \
  --name mysql-test \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=myapp \
  -p 3306:3306 \
  mysql:8.0

# Wait for container to be ready
docker logs -f mysql-test

# Run Fakestack
fakestack schema.json
```

## PostgreSQL

### Overview

PostgreSQL is a powerful, advanced open-source relational database with strong standards compliance.

**Pros:**
- Advanced features (JSON, arrays, full-text search)
- ACID compliant
- Extensible
- Great for complex queries

**Cons:**
- Requires server setup
- Can be resource intensive
- Steeper learning curve

### Configuration

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

### Connection Options

| Option | Required | Default | Description |
|--------|----------|---------|-------------|
| `dbtype` | Yes | - | Must be `"postgresql"` |
| `drivername` | Yes | - | Must be `"postgresql+psycopg2"` |
| `username` | Yes | - | PostgreSQL username |
| `password` | Yes | - | PostgreSQL password |
| `host` | Yes | `"localhost"` | Server host |
| `port` | No | `5432` | Server port |
| `database` | Yes | - | Database name |

### Prerequisites

1. **Install PostgreSQL:**
   ```bash
   # macOS (Homebrew)
   brew install postgresql
   brew services start postgresql
   
   # Ubuntu/Debian
   sudo apt-get install postgresql postgresql-contrib
   sudo systemctl start postgresql
   
   # Windows
   # Download from: https://www.postgresql.org/download/windows/
   ```

2. **Create Database:**
   ```bash
   # Connect as postgres user
   sudo -u postgres psql
   
   # Create database
   CREATE DATABASE myapp;
   
   # Create user (optional)
   CREATE USER myuser WITH PASSWORD 'mypassword';
   GRANT ALL PRIVILEGES ON DATABASE myapp TO myuser;
   ```

### Connection Examples

**Local development:**
```json
{
  "database": {
    "dbtype": "postgresql",
    "drivername": "postgresql+psycopg2",
    "username": "postgres",
    "password": "postgres",
    "host": "localhost",
    "port": 5432,
    "database": "myapp_dev"
  }
}
```

**Remote server:**
```json
{
  "database": {
    "dbtype": "postgresql",
    "drivername": "postgresql+psycopg2",
    "username": "app_user",
    "password": "secure_password",
    "host": "postgres.example.com",
    "port": 5432,
    "database": "production_db"
  }
}
```

**Docker container:**
```json
{
  "database": {
    "dbtype": "postgresql",
    "drivername": "postgresql+psycopg2",
    "username": "postgres",
    "password": "postgres",
    "host": "127.0.0.1",
    "port": 5432,
    "database": "test_db"
  }
}
```

### Docker Setup

```bash
# Start PostgreSQL container
docker run -d \
  --name postgres-test \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=myapp \
  -p 5432:5432 \
  postgres:15

# Wait for container to be ready
docker logs -f postgres-test

# Run Fakestack
fakestack schema.json
```

## Comparison

| Feature | SQLite | MySQL | PostgreSQL |
|---------|--------|-------|------------|
| Setup Complexity | ✅ None | ⚠️ Medium | ⚠️ Medium |
| Performance | ✅ Fast (small data) | ✅ Fast | ✅ Fast |
| Concurrent Writes | ❌ Limited | ✅ Excellent | ✅ Excellent |
| Data Types | ⚠️ Basic | ✅ Good | ✅ Advanced |
| Storage | File | Server | Server |
| Best For | Development, Testing | Web Apps, APIs | Complex Apps, Analytics |
| Network Access | ❌ No | ✅ Yes | ✅ Yes |
| Resource Usage | ✅ Low | ⚠️ Medium | ⚠️ Medium-High |

## Choosing a Database

### Use SQLite when:
- Developing or testing locally
- Building a prototype
- Small to medium dataset (< 1M rows)
- Single-user or low-concurrency application
- No server setup required

### Use MySQL when:
- Building web applications
- High-concurrency read/write operations
- Need proven reliability at scale
- Team familiar with MySQL
- Standard RDBMS features sufficient

### Use PostgreSQL when:
- Need advanced features (JSON, arrays, full-text search)
- Complex queries and analytics
- Data integrity is critical
- Need extensibility
- Want standards compliance

## Connection String Format

For manual connections, here are the equivalent connection strings:

**SQLite:**
```
sqlite:///path/to/database.db
```

**MySQL:**
```
mysql+mysqlconnector://username:password@host:port/database
```

**PostgreSQL:**
```
postgresql+psycopg2://username:password@host:port/database
```

## Troubleshooting

### SQLite

**File permissions:**
```bash
chmod 644 database.db
```

**Database locked:**
- Close all connections
- Check for other processes using the file
- Use `PRAGMA busy_timeout`

### MySQL

**Connection refused:**
```bash
# Check if MySQL is running
brew services list  # macOS
systemctl status mysql  # Linux

# Check port
netstat -an | grep 3306
```

**Access denied:**
```sql
# Check user permissions
SELECT user, host FROM mysql.user;
SHOW GRANTS FOR 'username'@'localhost';
```

### PostgreSQL

**Connection refused:**
```bash
# Check if PostgreSQL is running
brew services list  # macOS
systemctl status postgresql  # Linux

# Check port
netstat -an | grep 5432
```

**Authentication failed:**
```bash
# Edit pg_hba.conf
# Change 'peer' to 'md5' for password auth
sudo nano /etc/postgresql/*/main/pg_hba.conf
sudo systemctl restart postgresql
```

## Performance Tips

1. **Use Indexes**: Add indexes on frequently queried columns
2. **Batch Inserts**: Fakestack uses batch inserts for optimal performance
3. **Connection Pooling**: For repeated operations, keep connections alive
4. **Appropriate Data Types**: Use the right column types for your data
5. **Database Tuning**: Configure your database for your workload

## See Also

- [Schema Reference](schema-reference.md) - Complete schema options
- [Getting Started](getting-started.md) - Quick start guide
- [Examples](examples.md) - Database-specific examples
