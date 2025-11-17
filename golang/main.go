package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Version is set during build via -ldflags
var Version = "1.0.1"

const exampleSchema = `{
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
        },
        {
          "name": "first_name",
          "type": {"name": "string", "args": {"length": 50}},
          "options": {}
        },
        {
          "name": "last_name",
          "type": {"name": "string", "args": {"length": 50}},
          "options": {}
        }
      ],
      "indexes": []
    }
  ],
  "populate": [
    {
      "name": "users",
      "count": 50,
      "fields": [
        {"name": "username", "generator": "user_name"},
        {"name": "email", "generator": "email"},
        {"name": "first_name", "generator": "first_name"},
        {"name": "last_name", "generator": "last_name"}
      ]
    }
  ]
}
`

func main() {
	createTable := flag.Bool("c", false, "create database tables")
	createTableLong := flag.Bool("create-table", false, "create database tables")
	
	populateData := flag.Bool("p", false, "populate tables with fake data")
	populateDataLong := flag.Bool("populate-data", false, "populate tables with fake data")
	
	schemaFile := flag.String("f", "", "path to schema JSON file")
	schemaFileLong := flag.String("file", "", "path to schema JSON file")
	
	downloadSchema := flag.String("d", "", "download example schema to specified path (use '.' for current directory)")
	downloadSchemaLong := flag.String("download-schema", "", "download example schema to specified path")
	
	versionFlag := flag.Bool("v", false, "show version information")
	versionFlagLong := flag.Bool("version", false, "show version information")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Fakestack - Full-Stack Fake Data Generator\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	
	// Handle version flag
	if *versionFlag || *versionFlagLong {
		fmt.Printf("fakestack version %s\n", Version)
		return
	}

	// Merge short and long flags
	create := *createTable || *createTableLong
	populate := *populateData || *populateDataLong
	file := *schemaFile
	if file == "" {
		file = *schemaFileLong
	}
	download := *downloadSchema
	if download == "" {
		download = *downloadSchemaLong
	}

	// Handle download schema
	if download != "" {
		if err := downloadSchemaFile(download); err != nil {
			fmt.Fprintf(os.Stderr, "Error downloading schema: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Example schema downloaded successfully to schema.json")
		return
	}

	// Validate arguments
	if !create && !populate {
		fmt.Fprintf(os.Stderr, "Error: no arguments provided!!!\n")
		fmt.Fprintf(os.Stderr, "Use -c to create tables, -p to populate data, or -d to download example schema\n")
		flag.Usage()
		os.Exit(1)
	}

	if file == "" {
		fmt.Fprintf(os.Stderr, "Error: schema file is required\n")
		fmt.Fprintf(os.Stderr, "Use -f or --file to specify the schema file path\n")
		flag.Usage()
		os.Exit(1)
	}

	// Load schema
	schema, err := LoadSchema(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading schema: %v\n", err)
		os.Exit(1)
	}

	// Create database connection
	db, err := NewDatabase(schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Create tables
	if create {
		if err := db.DropTables(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to drop existing tables: %v\n", err)
		}

		if err := db.CreateTables(); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating tables: %v\n", err)
			os.Exit(1)
		}
	}

	// Populate data
	if populate {
		generator := NewGenerator(db)
		if err := db.PopulateData(generator); err != nil {
			fmt.Fprintf(os.Stderr, "Error populating data: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("✓ Operation completed successfully")
}

// downloadSchemaFile creates an example schema file
func downloadSchemaFile(path string) error {
	filename := "schema.json"
	if path != "." && path != "" {
		filename = filepath.Join(path, "schema.json")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, exampleSchema)
	return err
}
