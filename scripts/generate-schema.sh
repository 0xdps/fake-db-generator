#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}   Fakestack Schema Generator${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Ask for database type
echo -e "${YELLOW}Select database type:${NC}"
echo "1) SQLite"
echo "2) MySQL"
echo "3) PostgreSQL"
echo "4) MariaDB"
echo "5) MS SQL Server"
echo "6) CockroachDB"
read -p "Enter your choice (1-6): " db_choice

case $db_choice in
    1) dbtype="sqlite" ;;
    2) dbtype="mysql" ;;
    3) dbtype="postgres" ;;
    4) dbtype="mariadb" ;;
    5) dbtype="mssql" ;;
    6) dbtype="cockroachdb" ;;
    *) echo -e "${RED}Invalid choice!${NC}"; exit 1 ;;
esac

# Ask for database credentials (skip for SQLite)
if [ "$dbtype" != "sqlite" ]; then
    echo ""
    read -p "Enter database username (default: root): " username
    username=${username:-root}
    
    read -p "Enter database password (default: password): " password
    password=${password:-password}
    
    read -p "Enter database host (default: localhost): " host
    host=${host:-localhost}
    
    # Set default ports
    case $dbtype in
        mysql|mariadb) default_port="3306" ;;
        postgres) default_port="5432" ;;
        mssql) default_port="1433" ;;
        cockroachdb) default_port="26257" ;;
    esac
    
    read -p "Enter database port (default: $default_port): " port
    port=${port:-$default_port}
    host_with_port="${host}:${port}"
    
    read -p "Enter database name (default: testdb): " database
    database=${database:-testdb}
else
    database="test.db"
fi

# Ask for schema template
echo ""
echo -e "${YELLOW}Select schema template:${NC}"
echo "1) Users - Basic user management (id, username, email, created_at)"
echo "2) Employees - Employee records (id, first_name, last_name, email, department, salary, hire_date)"
echo "3) Products - E-commerce products (id, name, description, price, stock, category)"
echo "4) Orders - Order management (id, order_number, customer_email, total_amount, status, order_date)"
echo "5) Customers - Customer database (id, name, email, phone, address, city, country)"
echo "6) Blog Posts - Content management (id, title, content, author, published_at, views)"
echo "7) Inventory - Stock management (id, item_name, sku, quantity, location, last_updated)"
echo "8) Transactions - Financial records (id, transaction_id, amount, type, status, timestamp)"
echo "9) Students - Educational records (id, student_id, name, email, grade, enrollment_date)"
echo "10) Tasks - Task management (id, title, description, priority, status, due_date, assignee)"
read -p "Enter your choice (1-10): " schema_choice

# Ask for output filename
echo ""
read -p "Enter output filename (default: schema.json): " filename
filename=${filename:-schema.json}

# Generate schema based on choice
case $schema_choice in
    1) # Users
        table_name="users"
        row_count=50
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
EOF
)
        ;;
    
    2) # Employees
        table_name="employees"
        row_count=100
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
        {"name": "department", "generator": "random_string", "args": {"values": ["Engineering", "Sales", "Marketing", "HR", "Finance", "Operations"]}},
        {"name": "salary", "generator": "float", "args": {"min": 40000, "max": 150000}},
        {"name": "hire_date", "generator": "past_date"}
      ]
    }
  ]
}
EOF
)
        ;;
    
    3) # Products
        table_name="products"
        row_count=200
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
        {"name": "name", "generator": "product_name"},
        {"name": "description", "generator": "sentence", "args": {"word_count": 10}},
        {"name": "price", "generator": "float", "args": {"min": 9.99, "max": 999.99}},
        {"name": "stock", "generator": "integer", "args": {"min": 0, "max": 500}},
        {"name": "category", "generator": "random_string", "args": {"values": ["Electronics", "Clothing", "Books", "Home & Garden", "Toys", "Sports"]}}
      ]
    }
  ]
}
EOF
)
        ;;
    
    4) # Orders
        table_name="orders"
        row_count=500
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
  "tables": [
    {
      "name": "orders",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "order_number", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false, "unique": true}},
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
        {"name": "status", "generator": "random_string", "args": {"values": ["pending", "processing", "shipped", "delivered", "cancelled"]}},
        {"name": "order_date", "generator": "past_date"}
      ]
    }
  ]
}
EOF
)
        ;;
    
    5) # Customers
        table_name="customers"
        row_count=150
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
EOF
)
        ;;
    
    6) # Blog Posts
        table_name="posts"
        row_count=100
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
EOF
)
        ;;
    
    7) # Inventory
        table_name="inventory"
        row_count=300
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
        {"name": "item_name", "generator": "product_name"},
        {"name": "sku", "generator": "uuid"},
        {"name": "quantity", "generator": "integer", "args": {"min": 0, "max": 1000}},
        {"name": "location", "generator": "random_string", "args": {"values": ["Warehouse A", "Warehouse B", "Store 1", "Store 2", "Store 3"]}},
        {"name": "last_updated", "generator": "past_date"}
      ]
    }
  ]
}
EOF
)
        ;;
    
    8) # Transactions
        table_name="transactions"
        row_count=1000
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
        {"name": "type", "generator": "random_string", "args": {"values": ["credit", "debit", "transfer", "refund"]}},
        {"name": "status", "generator": "random_string", "args": {"values": ["completed", "pending", "failed", "cancelled"]}},
        {"name": "timestamp", "generator": "past_date"}
      ]
    }
  ]
}
EOF
)
        ;;
    
    9) # Students
        table_name="students"
        row_count=200
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
        {"name": "grade", "generator": "random_string", "args": {"values": ["A", "B", "C", "D", "F"]}},
        {"name": "enrollment_date", "generator": "past_date"}
      ]
    }
  ]
}
EOF
)
        ;;
    
    10) # Tasks
        table_name="tasks"
        row_count=150
        schema_json=$(cat <<EOF
{
  "database": {
    "dbtype": "$dbtype",
EOF
)
        if [ "$dbtype" != "sqlite" ]; then
            schema_json+=$(cat <<EOF

    "username": "$username",
    "password": "$password",
    "host": "$host_with_port",
    "database": "$database"
EOF
)
        else
            schema_json+=$(cat <<EOF

    "database": "$database"
EOF
)
        fi
        schema_json+=$(cat <<'EOF'

  },
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
        {"name": "priority", "generator": "random_string", "args": {"values": ["low", "medium", "high", "urgent"]}},
        {"name": "status", "generator": "random_string", "args": {"values": ["todo", "in_progress", "review", "done"]}},
        {"name": "due_date", "generator": "future_date"},
        {"name": "assignee", "generator": "name"}
      ]
    }
  ]
}
EOF
)
        ;;
    
    *) echo -e "${RED}Invalid choice!${NC}"; exit 1 ;;
esac

# Write to file
echo "$schema_json" > "$filename"

echo ""
echo -e "${GREEN}✓ Schema file generated successfully!${NC}"
echo -e "${BLUE}File: ${NC}$filename"
echo -e "${BLUE}Database: ${NC}$dbtype"
echo -e "${BLUE}Table: ${NC}$table_name"
echo -e "${BLUE}Rows: ${NC}$row_count"
echo ""
echo -e "${YELLOW}Next steps:${NC}"
echo -e "  1. Create tables: ${GREEN}fakestack -c -f $filename${NC}"
echo -e "  2. Populate data: ${GREEN}fakestack -p -f $filename${NC}"
echo -e "  3. Or do both:    ${GREEN}fakestack -c -p -f $filename${NC}"
echo ""
