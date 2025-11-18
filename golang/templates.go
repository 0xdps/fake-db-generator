package main

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

// SchemaTemplate represents a pre-built schema template
type SchemaTemplate struct {
	Name        string
	Description string
	TableName   string
	RowCount    int
	Generator   func(dbConfig DatabaseConfig) string
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	DBType   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// RunInteractiveGenerator runs the interactive schema generator
func RunInteractiveGenerator(outputFile string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========================================")
	fmt.Println("   Fakestack Schema Generator")
	fmt.Println("========================================")
	fmt.Println()

	// Ask for database type
	fmt.Println("Select database type:")
	fmt.Println("1) SQLite")
	fmt.Println("2) MySQL")
	fmt.Println("3) PostgreSQL")
	fmt.Println("4) MariaDB")
	fmt.Println("5) MS SQL Server")
	fmt.Println("6) CockroachDB")
	fmt.Print("Enter your choice (1-6): ")

	dbChoice, _ := reader.ReadString('\n')
	dbChoice = strings.TrimSpace(dbChoice)

	config := DatabaseConfig{}
	switch dbChoice {
	case "1":
		config.DBType = "sqlite"
	case "2":
		config.DBType = "mysql"
	case "3":
		config.DBType = "postgres"
	case "4":
		config.DBType = "mariadb"
	case "5":
		config.DBType = "mssql"
	case "6":
		config.DBType = "cockroachdb"
	default:
		return fmt.Errorf("invalid database choice")
	}

	// Ask for credentials (skip for SQLite)
	if config.DBType != "sqlite" {
		fmt.Println()
		fmt.Print("Enter database username (default: root): ")
		config.Username, _ = reader.ReadString('\n')
		config.Username = strings.TrimSpace(config.Username)
		if config.Username == "" {
			config.Username = "root"
		}

		fmt.Print("Enter database password (default: password): ")
		config.Password, _ = reader.ReadString('\n')
		config.Password = strings.TrimSpace(config.Password)
		if config.Password == "" {
			config.Password = "password"
		}

		fmt.Print("Enter database host (default: localhost): ")
		config.Host, _ = reader.ReadString('\n')
		config.Host = strings.TrimSpace(config.Host)
		if config.Host == "" {
			config.Host = "localhost"
		}

		// Set default ports
		defaultPort := "3306"
		switch config.DBType {
		case "mysql", "mariadb":
			defaultPort = "3306"
		case "postgres":
			defaultPort = "5432"
		case "mssql":
			defaultPort = "1433"
		case "cockroachdb":
			defaultPort = "26257"
		}

		fmt.Printf("Enter database port (default: %s): ", defaultPort)
		config.Port, _ = reader.ReadString('\n')
		config.Port = strings.TrimSpace(config.Port)
		if config.Port == "" {
			config.Port = defaultPort
		}

		fmt.Print("Enter database name (default: testdb): ")
		config.Database, _ = reader.ReadString('\n')
		config.Database = strings.TrimSpace(config.Database)
		if config.Database == "" {
			config.Database = "testdb"
		}
	} else {
		config.Database = "test.db"
	}

	// Ask for schema template
	fmt.Println()
	fmt.Println("Select schema template:")
	templates := GetSchemaTemplates()
	for i, tmpl := range templates {
		fmt.Printf("%d) %s - %s\n", i+1, tmpl.Name, tmpl.Description)
	}
	fmt.Printf("Enter your choice (1-%d): ", len(templates))

	schemaChoice, _ := reader.ReadString('\n')
	schemaChoice = strings.TrimSpace(schemaChoice)
	schemaIdx, err := strconv.Atoi(schemaChoice)
	if err != nil || schemaIdx < 1 || schemaIdx > len(templates) {
		return fmt.Errorf("invalid schema choice")
	}

	selectedTemplate := templates[schemaIdx-1]

	// Ask for output filename if not provided
	if outputFile == "" {
		fmt.Println()
		fmt.Print("Enter output filename (default: schema.json): ")
		outputFile, _ = reader.ReadString('\n')
		outputFile = strings.TrimSpace(outputFile)
		if outputFile == "" {
			outputFile = "schema.json"
		}
	}

	// Generate schema
	schemaContent := selectedTemplate.Generator(config)

	// Write to file
	if err := os.WriteFile(outputFile, []byte(schemaContent), 0644); err != nil {
		return fmt.Errorf("failed to write schema file: %w", err)
	}

	fmt.Println()
	fmt.Println("✓ Schema file generated successfully!")
	fmt.Printf("File: %s\n", outputFile)
	fmt.Printf("Database: %s\n", config.DBType)
	fmt.Printf("Table: %s\n", selectedTemplate.TableName)
	fmt.Printf("Rows: %d\n", selectedTemplate.RowCount)
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Printf("  1. Create tables: fakestack -c -f %s\n", outputFile)
	fmt.Printf("  2. Populate data: fakestack -p -f %s\n", outputFile)
	fmt.Printf("  3. Or do both:    fakestack -c -p -f %s\n", outputFile)
	fmt.Println()

	return nil
}

// GetSchemaTemplates returns all available schema templates
func GetSchemaTemplates() []SchemaTemplate {
	return []SchemaTemplate{
		{Name: "Users", Description: "Basic user management (id, username, email, created_at)", TableName: "users", RowCount: 50, Generator: generateUsersSchema},
		{Name: "Employees", Description: "Employee records (id, first_name, last_name, email, department, salary, hire_date)", TableName: "employees", RowCount: 100, Generator: generateEmployeesSchema},
		{Name: "Products", Description: "E-commerce products (id, name, description, price, stock, category)", TableName: "products", RowCount: 200, Generator: generateProductsSchema},
		{Name: "Orders", Description: "Order management (id, order_number, customer_email, total_amount, status, order_date)", TableName: "orders", RowCount: 500, Generator: generateOrdersSchema},
		{Name: "Customers", Description: "Customer database (id, name, email, phone, address, city, country)", TableName: "customers", RowCount: 150, Generator: generateCustomersSchema},
		{Name: "Blog Posts", Description: "Content management (id, title, content, author, published_at, views)", TableName: "posts", RowCount: 100, Generator: generateBlogPostsSchema},
		{Name: "Inventory", Description: "Stock management (id, item_name, sku, quantity, location, last_updated)", TableName: "inventory", RowCount: 300, Generator: generateInventorySchema},
		{Name: "Transactions", Description: "Financial records (id, transaction_id, amount, type, status, timestamp)", TableName: "transactions", RowCount: 1000, Generator: generateTransactionsSchema},
		{Name: "Students", Description: "Educational records (id, student_id, name, email, grade, enrollment_date)", TableName: "students", RowCount: 200, Generator: generateStudentsSchema},
		{Name: "Tasks", Description: "Task management (id, title, description, priority, status, due_date, assignee)", TableName: "tasks", RowCount: 150, Generator: generateTasksSchema},
	}
}

// Helper function to build database config JSON
func buildDatabaseConfigJSON(config DatabaseConfig) string {
	if config.DBType == "sqlite" {
		return fmt.Sprintf(`  "database": {
    "dbtype": "%s",
    "database": "%s"
  }`, config.DBType, config.Database)
	}
	hostWithPort := fmt.Sprintf("%s:%s", config.Host, config.Port)
	return fmt.Sprintf(`  "database": {
    "dbtype": "%s",
    "username": "%s",
    "password": "%s",
    "host": "%s",
    "database": "%s"
  }`, config.DBType, config.Username, config.Password, hostWithPort, config.Database)
}

// Schema generator functions
func generateUsersSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "users",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "username", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}},
        {"name": "email", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false, "unique": true}},
        {"name": "created_at", "type": "datetime", "options": {"nullable": false}}
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
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateEmployeesSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "employees",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "first_name", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false}},
        {"name": "last_name", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false}},
        {"name": "email", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false, "unique": true}},
        {"name": "department", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false}},
        {"name": "salary", "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}}, "options": {"nullable": false}},
        {"name": "hire_date", "type": "date", "options": {"nullable": false}}
      ]
    }
  ],
  "populate": [
    {
      "name": "employees",
      "count": 100,
      "fields": [
        {"name": "first_name", "generator": "first_name"},
        {"name": "last_name", "generator": "last_name"},
        {"name": "email", "generator": "email"},
        {"name": "department", "generator": "random_from", "args": ["Engineering", "Sales", "Marketing", "HR", "Finance", "Operations"]},
        {"name": "salary", "generator": "float", "args": {"min": 40000, "max": 150000}},
        {"name": "hire_date", "generator": "past_date"}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateProductsSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "products",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "name", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false}},
        {"name": "description", "type": "text", "options": {"nullable": true}},
        {"name": "price", "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}}, "options": {"nullable": false}},
        {"name": "stock", "type": "integer", "options": {"nullable": false}},
        {"name": "category", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false}}
      ]
    }
  ],
  "populate": [
    {
      "name": "products",
      "count": 200,
      "fields": [
        {"name": "name", "generator": "word"},
        {"name": "description", "generator": "sentence", "args": {"word_count": 10}},
        {"name": "price", "generator": "float", "args": {"min": 9.99, "max": 999.99}},
        {"name": "stock", "generator": "integer", "args": {"min": 0, "max": 500}},
        {"name": "category", "generator": "random_from", "args": ["Electronics", "Clothing", "Books", "Home & Garden", "Toys", "Sports"]}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateOrdersSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "orders",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "order_number", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}},
        {"name": "customer_email", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false}},
        {"name": "total_amount", "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}}, "options": {"nullable": false}},
        {"name": "status", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}},
        {"name": "order_date", "type": "datetime", "options": {"nullable": false}}
      ]
    }
  ],
  "populate": [
    {
      "name": "orders",
      "count": 500,
      "fields": [
        {"name": "order_number", "generator": "uuid"},
        {"name": "customer_email", "generator": "email"},
        {"name": "total_amount", "generator": "float", "args": {"min": 10.00, "max": 1000.00}},
        {"name": "status", "generator": "random_from", "args": ["pending", "processing", "shipped", "delivered", "cancelled"]},
        {"name": "order_date", "generator": "past_date"}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateCustomersSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "customers",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "name", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false}},
        {"name": "email", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false, "unique": true}},
        {"name": "phone", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": true}},
        {"name": "address", "type": {"name": "string", "args": {"length": 200}}, "options": {"nullable": true}},
        {"name": "city", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": true}},
        {"name": "country", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false}}
      ]
    }
  ],
  "populate": [
    {
      "name": "customers",
      "count": 150,
      "fields": [
        {"name": "name", "generator": "name"},
        {"name": "email", "generator": "email"},
        {"name": "phone", "generator": "phone_number"},
        {"name": "address", "generator": "address"},
        {"name": "city", "generator": "city"},
        {"name": "country", "generator": "country"}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateBlogPostsSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "posts",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "title", "type": {"name": "string", "args": {"length": 200}}, "options": {"nullable": false}},
        {"name": "content", "type": "text", "options": {"nullable": false}},
        {"name": "author", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false}},
        {"name": "published_at", "type": "datetime", "options": {"nullable": false}},
        {"name": "views", "type": "integer", "options": {"nullable": false}}
      ]
    }
  ],
  "populate": [
    {
      "name": "posts",
      "count": 100,
      "fields": [
        {"name": "title", "generator": "sentence", "args": {"word_count": 6}},
        {"name": "content", "generator": "paragraph", "args": {"sentence_count": 5}},
        {"name": "author", "generator": "name"},
        {"name": "published_at", "generator": "past_date"},
        {"name": "views", "generator": "integer", "args": {"min": 0, "max": 10000}}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateInventorySchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "inventory",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "item_name", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false}},
        {"name": "sku", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}},
        {"name": "quantity", "type": "integer", "options": {"nullable": false}},
        {"name": "location", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false}},
        {"name": "last_updated", "type": "datetime", "options": {"nullable": false}}
      ]
    }
  ],
  "populate": [
    {
      "name": "inventory",
      "count": 300,
      "fields": [
        {"name": "item_name", "generator": "word"},
        {"name": "sku", "generator": "uuid"},
        {"name": "quantity", "generator": "integer", "args": {"min": 0, "max": 1000}},
        {"name": "location", "generator": "random_from", "args": ["Warehouse A", "Warehouse B", "Store 1", "Store 2", "Store 3"]},
        {"name": "last_updated", "generator": "past_date"}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateTransactionsSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "transactions",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "transaction_id", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}},
        {"name": "amount", "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}}, "options": {"nullable": false}},
        {"name": "type", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}},
        {"name": "status", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}},
        {"name": "timestamp", "type": "datetime", "options": {"nullable": false}}
      ]
    }
  ],
  "populate": [
    {
      "name": "transactions",
      "count": 1000,
      "fields": [
        {"name": "transaction_id", "generator": "uuid"},
        {"name": "amount", "generator": "float", "args": {"min": 1.00, "max": 10000.00}},
        {"name": "type", "generator": "random_from", "args": ["credit", "debit", "transfer", "refund"]},
        {"name": "status", "generator": "random_from", "args": ["completed", "pending", "failed", "cancelled"]},
        {"name": "timestamp", "generator": "past_date"}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateStudentsSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "students",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "student_id", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false, "unique": true}},
        {"name": "name", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false}},
        {"name": "email", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false, "unique": true}},
        {"name": "grade", "type": {"name": "string", "args": {"length": 10}}, "options": {"nullable": false}},
        {"name": "enrollment_date", "type": "date", "options": {"nullable": false}}
      ]
    }
  ],
  "populate": [
    {
      "name": "students",
      "count": 200,
      "fields": [
        {"name": "student_id", "generator": "uuid"},
        {"name": "name", "generator": "name"},
        {"name": "email", "generator": "email"},
        {"name": "grade", "generator": "random_from", "args": ["A", "B", "C", "D", "F"]},
        {"name": "enrollment_date", "generator": "past_date"}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}

func generateTasksSchema(config DatabaseConfig) string {
	return fmt.Sprintf(`{
%s,
  "tables": [
    {
      "name": "tasks",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "title", "type": {"name": "string", "args": {"length": 200}}, "options": {"nullable": false}},
        {"name": "description", "type": "text", "options": {"nullable": true}},
        {"name": "priority", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}},
        {"name": "status", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}},
        {"name": "due_date", "type": "date", "options": {"nullable": true}},
        {"name": "assignee", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": true}}
      ]
    }
  ],
  "populate": [
    {
      "name": "tasks",
      "count": 150,
      "fields": [
        {"name": "title", "generator": "sentence", "args": {"word_count": 5}},
        {"name": "description", "generator": "paragraph", "args": {"sentence_count": 3}},
        {"name": "priority", "generator": "random_from", "args": ["low", "medium", "high", "urgent"]},
        {"name": "status", "generator": "random_from", "args": ["todo", "in_progress", "review", "done"]},
        {"name": "due_date", "generator": "future_date"},
        {"name": "assignee", "generator": "name"}
      ]
    }
  ]
}
`, buildDatabaseConfigJSON(config))
}
