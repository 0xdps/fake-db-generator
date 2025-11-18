package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// DbType represents the database type
type DbType string

const (
	MySQL       DbType = "mysql"
	Postgres    DbType = "postgres"
	PostgresAlt DbType = "psql" // Alternative name for backwards compatibility
	SQLite      DbType = "sqlite"
	MSSQL       DbType = "mssql"
	MariaDB     DbType = "mariadb"
	CockroachDB DbType = "cockroachdb"
)

// DbOptions represents database connection configuration
type DbOptions struct {
	DbType     DbType `json:"dbtype"`
	DriverName string `json:"drivername"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Host       string `json:"host,omitempty"`
	Database   string `json:"database"`
	Echo       bool   `json:"echo,omitempty"`
}

// ColumnType represents a database column type
type ColumnType struct {
	Name string                 `json:"name"`
	Args map[string]interface{} `json:"args,omitempty"`
}

// TableColumn represents a database column
type TableColumn struct {
	Name    string                 `json:"name"`
	Type    interface{}            `json:"type"` // Can be string or ColumnType
	Options map[string]interface{} `json:"options"`
}

// ParsedTableColumn is the processed version of TableColumn
type ParsedTableColumn struct {
	Name    string
	Type    ColumnType
	Options map[string]interface{}
}

// TableIndex represents a database index
type TableIndex struct {
	Name    string   `json:"name"`
	Columns []string `json:"columns"`
}

// DbTable represents a database table
type DbTable struct {
	Name    string        `json:"name"`
	Columns []TableColumn `json:"columns"`
	Indexes []TableIndex  `json:"indexes,omitempty"`
}

// PopulateField represents a field to populate with fake data
type PopulateField struct {
	Name      string      `json:"name"`
	Generator string      `json:"generator"`
	Args      interface{} `json:"args,omitempty"`
}

// IsDbAccess checks if the generator needs database access
func (f *PopulateField) IsDbAccess() bool {
	return f.Generator == "db_random_item"
}

// DbPopulate represents data population configuration
type DbPopulate struct {
	Name   string          `json:"name"`
	Fields []PopulateField `json:"fields"`
	Count  int             `json:"count"`
}

// DbSchema represents the complete database schema
type DbSchema struct {
	Database DbOptions    `json:"database"`
	Tables   []DbTable    `json:"tables,omitempty"`
	Populate []DbPopulate `json:"populate,omitempty"`
}

// LoadSchema loads a schema from a JSON file
func LoadSchema(filename string) (*DbSchema, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	var schema DbSchema
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
	}

	return &schema, nil
}

// ParseColumnType converts interface{} type to ColumnType
func ParseColumnType(t interface{}) ColumnType {
	switch v := t.(type) {
	case string:
		return ColumnType{Name: v}
	case map[string]interface{}:
		ct := ColumnType{}
		if name, ok := v["name"].(string); ok {
			ct.Name = name
		}
		if args, ok := v["args"].(map[string]interface{}); ok {
			ct.Args = args
		}
		return ct
	default:
		return ColumnType{Name: "string"}
	}
}

// ParseTableColumn converts TableColumn to ParsedTableColumn
func ParseTableColumn(col TableColumn) ParsedTableColumn {
	return ParsedTableColumn{
		Name:    col.Name,
		Type:    ParseColumnType(col.Type),
		Options: col.Options,
	}
}
