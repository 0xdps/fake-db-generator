package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Database represents a database connection
type Database struct {
	conn   *sql.DB
	schema *DbSchema
	driver string
}

// NewDatabase creates a new database connection
func NewDatabase(schema *DbSchema) (*Database, error) {
	connStr := buildConnectionString(schema.Database)
	driver := getDriver(schema.Database.DbType)

	conn, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{
		conn:   conn,
		schema: schema,
		driver: driver,
	}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.conn.Close()
}

// getDriver returns the SQL driver name
func getDriver(dbType DbType) string {
	switch dbType {
	case MySQL, MariaDB:
		return "mysql"
	case Postgres, PostgresAlt, CockroachDB:
		return "postgres"
	case SQLite:
		return "sqlite3"
	case MSSQL:
		return "sqlserver"
	default:
		return "sqlite3"
	}
}

// buildConnectionString builds a database connection string
func buildConnectionString(opts DbOptions) string {
	switch opts.DbType {
	case MySQL, MariaDB:
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			opts.Username, opts.Password, opts.Host, opts.Database)
	case Postgres, PostgresAlt, CockroachDB:
		return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			opts.Username, opts.Password, opts.Host, opts.Database)
	case SQLite:
		return opts.Database
	case MSSQL:
		return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
			opts.Username, opts.Password, opts.Host, opts.Database)
	default:
		return opts.Database
	}
}

// CreateTables creates all tables defined in the schema
func (db *Database) CreateTables() error {
	for _, table := range db.schema.Tables {
		if err := db.createTable(table); err != nil {
			return fmt.Errorf("failed to create table %s: %w", table.Name, err)
		}
		fmt.Printf("Table %s successfully created\n", table.Name)
	}
	return nil
}

// DropTables drops all tables defined in the schema
func (db *Database) DropTables() error {
	for _, table := range db.schema.Tables {
		var query string
		if db.driver == "sqlserver" {
			query = fmt.Sprintf("IF OBJECT_ID('[%s]', 'U') IS NOT NULL DROP TABLE [%s]", table.Name, table.Name)
		} else {
			query = fmt.Sprintf("DROP TABLE IF EXISTS %s", table.Name)
		}
		if _, err := db.conn.Exec(query); err != nil {
			return fmt.Errorf("failed to drop table %s: %w", table.Name, err)
		}
	}
	return nil
}

// createTable creates a single table
func (db *Database) createTable(table DbTable) error {
	var columns []string

	for _, col := range table.Columns {
		parsed := ParseTableColumn(col)
		colDef := db.buildColumnDefinition(parsed)
		columns = append(columns, colDef)
	}

	var query string
	if db.driver == "sqlserver" {
		query = fmt.Sprintf("IF OBJECT_ID('[%s]', 'U') IS NULL CREATE TABLE [%s] (\n  %s\n)",
			table.Name, table.Name, strings.Join(columns, ",\n  "))
	} else {
		query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n  %s\n)",
			table.Name, strings.Join(columns, ",\n  "))
	}

	_, err := db.conn.Exec(query)
	return err
}

// buildColumnDefinition builds a SQL column definition
func (db *Database) buildColumnDefinition(col ParsedTableColumn) string {
	sqlType := db.getSQLType(col.Type)
	colName := col.Name
	if db.driver == "sqlserver" {
		colName = fmt.Sprintf("[%s]", col.Name)
	}
	def := fmt.Sprintf("%s %s", colName, sqlType)

	// Handle options
	if primary, ok := col.Options["primary_key"].(bool); ok && primary {
		def += " PRIMARY KEY"
		if autoInc, ok := col.Options["autoincrement"].(bool); ok && autoInc {
			if db.driver == "sqlite3" {
				def += " AUTOINCREMENT"
			} else if db.driver == "mysql" {
				def += " AUTO_INCREMENT"
			} else if db.driver == "sqlserver" {
				def += " IDENTITY(1,1)"
			} else {
				// PostgreSQL and CockroachDB use SERIAL or BIGSERIAL
				def = strings.Replace(def, sqlType, "SERIAL", 1)
			}
		}
	}

	if nullable, ok := col.Options["nullable"].(bool); ok && !nullable {
		def += " NOT NULL"
	}

	if unique, ok := col.Options["unique"].(bool); ok && unique {
		def += " UNIQUE"
	}

	return def
}

// getSQLType converts schema type to SQL type
func (db *Database) getSQLType(colType ColumnType) string {
	switch colType.Name {
	case "integer":
		return "INTEGER"
	case "number":
		precision := 10
		scale := 2
		if colType.Args != nil {
			if p, ok := colType.Args["precision"].(float64); ok {
				precision = int(p)
			}
			if s, ok := colType.Args["scale"].(float64); ok {
				scale = int(s)
			}
		}
		return fmt.Sprintf("NUMERIC(%d, %d)", precision, scale)
	case "string":
		length := 255
		if colType.Args != nil {
			if l, ok := colType.Args["length"].(float64); ok {
				length = int(l)
			}
		}
		return fmt.Sprintf("VARCHAR(%d)", length)
	case "date":
		return "DATE"
	case "datetime":
		return "DATETIME"
	case "text":
		return "TEXT"
	case "float":
		// Use REAL for SQLite, FLOAT for others
		if db.driver == "sqlite3" {
			return "REAL"
		}
		return "FLOAT"
	case "decimal":
		precision := 10
		scale := 2
		if colType.Args != nil {
			if p, ok := colType.Args["precision"].(float64); ok {
				precision = int(p)
			}
			if s, ok := colType.Args["scale"].(float64); ok {
				scale = int(s)
			}
		}
		return fmt.Sprintf("DECIMAL(%d, %d)", precision, scale)
	case "boolean":
		// MSSQL uses BIT instead of BOOLEAN
		if db.driver == "sqlserver" {
			return "BIT"
		}
		return "BOOLEAN"
	default:
		return "VARCHAR(255)"
	}
}

// Insert inserts a row into a table
func (db *Database) Insert(tableName string, data map[string]interface{}) error {
	var columns []string
	var placeholders []string
	var values []interface{}
	i := 1

	for col, val := range data {
		if db.driver == "sqlserver" {
			columns = append(columns, fmt.Sprintf("[%s]", col))
			placeholders = append(placeholders, fmt.Sprintf("@p%d", i))
		} else {
			columns = append(columns, col)
			if db.driver == "postgres" {
				placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			} else {
				placeholders = append(placeholders, "?")
			}
		}
		values = append(values, val)
		i++
	}

	var tableName_quoted string
	if db.driver == "sqlserver" {
		tableName_quoted = fmt.Sprintf("[%s]", tableName)
	} else {
		tableName_quoted = tableName
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName_quoted,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	_, err := db.conn.Exec(query, values...)
	return err
}

// GetRandomValue gets a random value from a column
func (db *Database) GetRandomValue(tableName, columnName string) (interface{}, error) {
	var query string
	tableQuoted := tableName
	columnQuoted := columnName
	if db.driver == "sqlserver" {
		tableQuoted = fmt.Sprintf("[%s]", tableName)
		columnQuoted = fmt.Sprintf("[%s]", columnName)
	}

	switch db.driver {
	case "sqlite3", "postgres":
		query = fmt.Sprintf("SELECT %s FROM %s ORDER BY RANDOM() LIMIT 1", columnQuoted, tableQuoted)
	case "mysql":
		query = fmt.Sprintf("SELECT %s FROM %s ORDER BY RAND() LIMIT 1", columnQuoted, tableQuoted)
	case "sqlserver":
		query = fmt.Sprintf("SELECT TOP 1 %s FROM %s ORDER BY NEWID()", columnQuoted, tableQuoted)
	}

	var value interface{}
	err := db.conn.QueryRow(query).Scan(&value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
