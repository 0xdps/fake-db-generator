# Database Support

Fakestack supports **6 major database systems** with native drivers and optimized connections:

- **SQLite** - Lightweight, file-based database
- **MySQL** - Popular open-source relational database
- **PostgreSQL** - Advanced open-source relational database
- **MariaDB** - MySQL-compatible database with enhanced features
- **MS SQL Server** - Microsoft's enterprise database system
- **CockroachDB** - Distributed SQL database (PostgreSQL-compatible)

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

## MariaDB

### Overview

MariaDB is a MySQL-compatible database with enhanced performance, security features, and additional storage engines.

**Pros:**
- Drop-in MySQL replacement
- Better performance than MySQL in many workloads
- More storage engines (Aria, ColumnStore, etc.)
- Active open-source development

**Cons:**
- Some MySQL features not available
- Less widely adopted than MySQL

### Configuration

```json
{
  "database": {
    "dbtype": "mariadb",
    "username": "root",
    "password": "yourpassword",
    "host": "localhost:3306",
    "database": "testdb"
  }
}
```

### Connection Options

| Option | Required | Default | Description |
|--------|----------|---------|-------------|
| `dbtype` | Yes | - | Must be `"mariadb"` |
| `username` | Yes | - | Database user |
| `password` | Yes | - | User password |
| `host` | Yes | - | Host and port (e.g., `localhost:3306`) |
| `database` | Yes | - | Database name |

### Usage Example

```bash
# Create and populate
fakestack -c -p -f schema.json
```

**Schema example:**
```json
{
  "database": {
    "dbtype": "mariadb",
    "username": "root",
    "password": "password",
    "host": "localhost:3306",
    "database": "mydb"
  },
  "tables": [...],
  "populate": [...]
}
```

## MS SQL Server

### Overview

Microsoft SQL Server is an enterprise-grade relational database management system with advanced analytics and integration capabilities.

**Pros:**
- Enterprise features and support
- Excellent Windows integration
- Advanced analytics and BI tools
- Strong security features

**Cons:**
- Primarily Windows-focused (Linux support improving)
- Licensing costs for production
- More resource-intensive

### Configuration

```json
{
  "database": {
    "dbtype": "mssql",
    "username": "sa",
    "password": "YourPassword123!",
    "host": "localhost:1433",
    "database": "testdb"
  }
}
```

### Connection Options

| Option | Required | Default | Description |
|--------|----------|---------|-------------|
| `dbtype` | Yes | - | Must be `"mssql"` |
| `username` | Yes | - | Database user (typically `sa`) |
| `password` | Yes | - | Strong password required |
| `host` | Yes | - | Host and port (e.g., `localhost:1433`) |
| `database` | Yes | - | Database name |

### Docker Setup

```bash
docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=YourPassword123!' \
  -p 1433:1433 --name mssql \
  mcr.microsoft.com/mssql/server:2022-latest
```

### Usage Example

```bash
fakestack -c -p -f mssql-schema.json
```

### Special Considerations

- **Password requirements**: Must be strong (uppercase, lowercase, numbers, special chars)
- **IDENTITY columns**: Use `autoincrement: true` for auto-incrementing primary keys
- **Connection timeout**: May need longer timeout for first connection

## CockroachDB

### Overview

CockroachDB is a distributed SQL database that is PostgreSQL-compatible and designed for cloud-native applications.

**Pros:**
- Horizontal scalability
- Built-in replication and consistency
- PostgreSQL wire protocol compatibility
- Resilient to node failures

**Cons:**
- More complex setup than traditional databases
- Different performance characteristics
- Some PostgreSQL features not supported

### Configuration

```json
{
  "database": {
    "dbtype": "cockroachdb",
    "username": "root",
    "password": "",
    "host": "localhost:26257",
    "database": "testdb"
  }
}
```

### Connection Options

| Option | Required | Default | Description |
|--------|----------|---------|-------------|
| `dbtype` | Yes | - | Must be `"cockroachdb"` |
| `username` | Yes | - | Database user (default: `root`) |
| `password` | No | `""` | Password (empty for insecure mode) |
| `host` | Yes | - | Host and port (default: `26257`) |
| `database` | Yes | - | Database name |

### Docker Setup

```bash
# Start single-node cluster (insecure for testing)
docker run -d -p 26257:26257 -p 8080:8080 \
  --name cockroach \
  cockroachdb/cockroach:latest start-single-node --insecure

# Create database
docker exec -it cockroach ./cockroach sql --insecure \
  --execute="CREATE DATABASE testdb;"
```

### Usage Example

```bash
fakestack -c -p -f cockroachdb-schema.json
```

### Special Considerations

- Uses PostgreSQL driver internally
- Default port is `26257` (not `5432`)
- Supports most PostgreSQL SQL syntax
- Better suited for distributed deployments

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
