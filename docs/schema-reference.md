# Schema Reference

Complete reference for creating Fakestack JSON schemas.

## Schema Structure

A Fakestack schema consists of three main sections:

```json
{
  "database": { ... },    // Database connection configuration
  "tables": [ ... ],      // Table definitions
  "populate": [ ... ]     // Data population instructions
}
```

## Database Configuration

### SQLite

```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "path/to/database.db"
  }
}
```

**Options:**
- `dbtype`: Must be `"sqlite"`
- `drivername`: Must be `"sqlite"`
- `database`: Path to SQLite database file (will be created if it doesn't exist)

### MySQL

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

**Options:**
- `dbtype`: Must be `"mysql"`
- `drivername`: Must be `"mysql+mysqlconnector"`
- `username`: MySQL username
- `password`: MySQL password
- `host`: MySQL server host (default: `"localhost"`)
- `port`: MySQL server port (default: `3306`)
- `database`: Database name (must exist)

### PostgreSQL

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

**Options:**
- `dbtype`: Must be `"postgresql"`
- `drivername`: Must be `"postgresql+psycopg2"`
- `username`: PostgreSQL username
- `password`: PostgreSQL password
- `host`: PostgreSQL server host (default: `"localhost"`)
- `port`: PostgreSQL server port (default: `5432`)
- `database`: Database name (must exist)

## Table Definitions

### Basic Table

```json
{
  "name": "users",
  "columns": [
    {
      "name": "id",
      "type": "integer",
      "options": {
        "primary_key": true,
        "autoincrement": true
      }
    },
    {
      "name": "username",
      "type": {
        "name": "string",
        "args": {
          "length": 50
        }
      },
      "options": {
        "nullable": false,
        "unique": true
      }
    }
  ],
  "indexes": []
}
```

### Column Types

#### Simple Types

```json
// Integer
{ "name": "age", "type": "integer", "options": {} }

// String (requires length)
{ "name": "name", "type": {"name": "string", "args": {"length": 100}}, "options": {} }

// Text (no length limit)
{ "name": "description", "type": "text", "options": {} }

// Boolean
{ "name": "is_active", "type": "boolean", "options": {} }

// Date
{ "name": "birth_date", "type": "date", "options": {} }

// DateTime
{ "name": "created_at", "type": "datetime", "options": {} }

// Time
{ "name": "login_time", "type": "time", "options": {} }

// Float
{ "name": "price", "type": "float", "options": {} }

// Decimal (requires precision and scale)
{ 
  "name": "amount", 
  "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}}, 
  "options": {} 
}
```

### Column Options

```json
{
  "options": {
    "primary_key": true,      // Mark as primary key
    "autoincrement": true,    // Auto-increment (for integers)
    "nullable": false,        // Allow NULL values
    "unique": true,           // Enforce unique constraint
    "default": "value"        // Default value
  }
}
```

### Foreign Keys

```json
{
  "name": "user_id",
  "type": "integer",
  "options": {
    "nullable": false,
    "foreign_key": {
      "table": "users",
      "column": "id",
      "ondelete": "CASCADE"    // Optional: CASCADE, SET NULL, RESTRICT
    }
  }
}
```

### Indexes

```json
{
  "name": "users",
  "columns": [...],
  "indexes": [
    {
      "name": "idx_username",
      "columns": ["username"],
      "unique": false
    },
    {
      "name": "idx_email_username",
      "columns": ["email", "username"],
      "unique": true
    }
  ]
}
```

## Data Population

### Basic Population

```json
{
  "populate": [
    {
      "name": "users",
      "count": 100,
      "fields": [
        { "name": "username", "generator": "user_name" },
        { "name": "email", "generator": "email" },
        { "name": "created_at", "generator": "past_date" }
      ]
    }
  ]
}
```

### Generator Arguments

```json
{
  "name": "age",
  "generator": "random_int",
  "args": {
    "min": 18,
    "max": 65
  }
}
```

### Random From List

```json
{
  "name": "status",
  "generator": "random_from",
  "args": {
    "choices": ["active", "inactive", "pending"]
  }
}
```

### Foreign Key References

When populating tables with foreign keys, populate the parent table first:

```json
{
  "populate": [
    {
      "name": "users",
      "count": 50,
      "fields": [...]
    },
    {
      "name": "posts",
      "count": 200,
      "fields": [
        { "name": "title", "generator": "sentence" },
        { "name": "user_id", "generator": "foreign_key", "args": {"table": "users"} }
      ]
    }
  ]
}
```

## Complete Example

```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "blog.db"
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
        },
        {
          "name": "created_at",
          "type": "datetime",
          "options": {"nullable": false}
        }
      ],
      "indexes": [
        {"name": "idx_username", "columns": ["username"], "unique": true}
      ]
    },
    {
      "name": "posts",
      "columns": [
        {
          "name": "id",
          "type": "integer",
          "options": {"primary_key": true, "autoincrement": true}
        },
        {
          "name": "user_id",
          "type": "integer",
          "options": {
            "nullable": false,
            "foreign_key": {"table": "users", "column": "id", "ondelete": "CASCADE"}
          }
        },
        {
          "name": "title",
          "type": {"name": "string", "args": {"length": 200}},
          "options": {"nullable": false}
        },
        {
          "name": "content",
          "type": "text",
          "options": {}
        },
        {
          "name": "published_at",
          "type": "datetime",
          "options": {}
        }
      ],
      "indexes": [
        {"name": "idx_user_id", "columns": ["user_id"], "unique": false}
      ]
    }
  ],
  "populate": [
    {
      "name": "users",
      "count": 50,
      "fields": [
        {"name": "username", "generator": "user_name"},
        {"name": "email", "generator": "email"},
        {"name": "created_at", "generator": "past_date"}
      ]
    },
    {
      "name": "posts",
      "count": 200,
      "fields": [
        {"name": "user_id", "generator": "foreign_key", "args": {"table": "users"}},
        {"name": "title", "generator": "sentence"},
        {"name": "content", "generator": "paragraph"},
        {"name": "published_at", "generator": "past_date"}
      ]
    }
  ]
}
```

## Best Practices

1. **Order Matters**: Define tables in order of dependencies (parent tables before child tables)
2. **Populate in Order**: Populate parent tables before child tables with foreign keys
3. **Use Indexes**: Add indexes on frequently queried columns
4. **Unique Constraints**: Use unique constraints for data integrity
5. **Sensible Counts**: Start with smaller counts for testing, scale up later
6. **Foreign Keys**: Always validate foreign key relationships match your data model

## See Also

- [Data Generators](generators.md) - All available generators
- [Database Support](databases.md) - Database-specific details
- [Examples](examples.md) - More complex schemas
