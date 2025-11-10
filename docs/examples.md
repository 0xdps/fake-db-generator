# Examples

Real-world schema examples for common use cases.

## Table of Contents

- [E-Commerce Application](#e-commerce-application)
- [Blog Platform](#blog-platform)
- [Social Network](#social-network)
- [Task Management](#task-management)
- [Event Registration](#event-registration)

---

## E-Commerce Application

Complete schema for an online store with products, customers, and orders.

```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "ecommerce.db"
  },
  "tables": [
    {
      "name": "customers",
      "columns": [
        {
          "name": "id",
          "type": "integer",
          "options": {"primary_key": true, "autoincrement": true}
        },
        {
          "name": "first_name",
          "type": {"name": "string", "args": {"length": 50}},
          "options": {"nullable": false}
        },
        {
          "name": "last_name",
          "type": {"name": "string", "args": {"length": 50}},
          "options": {"nullable": false}
        },
        {
          "name": "email",
          "type": {"name": "string", "args": {"length": 100}},
          "options": {"nullable": false, "unique": true}
        },
        {
          "name": "phone",
          "type": {"name": "string", "args": {"length": 20}},
          "options": {}
        },
        {
          "name": "address",
          "type": "text",
          "options": {}
        },
        {
          "name": "city",
          "type": {"name": "string", "args": {"length": 50}},
          "options": {}
        },
        {
          "name": "state",
          "type": {"name": "string", "args": {"length": 50}},
          "options": {}
        },
        {
          "name": "postcode",
          "type": {"name": "string", "args": {"length": 10}},
          "options": {}
        },
        {
          "name": "created_at",
          "type": "datetime",
          "options": {"nullable": false}
        }
      ],
      "indexes": [
        {"name": "idx_email", "columns": ["email"], "unique": true}
      ]
    },
    {
      "name": "categories",
      "columns": [
        {
          "name": "id",
          "type": "integer",
          "options": {"primary_key": true, "autoincrement": true}
        },
        {
          "name": "name",
          "type": {"name": "string", "args": {"length": 100}},
          "options": {"nullable": false, "unique": true}
        },
        {
          "name": "description",
          "type": "text",
          "options": {}
        }
      ],
      "indexes": []
    },
    {
      "name": "products",
      "columns": [
        {
          "name": "id",
          "type": "integer",
          "options": {"primary_key": true, "autoincrement": true}
        },
        {
          "name": "category_id",
          "type": "integer",
          "options": {
            "nullable": false,
            "foreign_key": {"table": "categories", "column": "id"}
          }
        },
        {
          "name": "name",
          "type": {"name": "string", "args": {"length": 200}},
          "options": {"nullable": false}
        },
        {
          "name": "description",
          "type": "text",
          "options": {}
        },
        {
          "name": "price",
          "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}},
          "options": {"nullable": false}
        },
        {
          "name": "stock",
          "type": "integer",
          "options": {"nullable": false}
        },
        {
          "name": "sku",
          "type": {"name": "string", "args": {"length": 50}},
          "options": {"nullable": false, "unique": true}
        },
        {
          "name": "created_at",
          "type": "datetime",
          "options": {"nullable": false}
        }
      ],
      "indexes": [
        {"name": "idx_category", "columns": ["category_id"], "unique": false},
        {"name": "idx_sku", "columns": ["sku"], "unique": true}
      ]
    },
    {
      "name": "orders",
      "columns": [
        {
          "name": "id",
          "type": "integer",
          "options": {"primary_key": true, "autoincrement": true}
        },
        {
          "name": "customer_id",
          "type": "integer",
          "options": {
            "nullable": false,
            "foreign_key": {"table": "customers", "column": "id", "ondelete": "CASCADE"}
          }
        },
        {
          "name": "order_date",
          "type": "datetime",
          "options": {"nullable": false}
        },
        {
          "name": "status",
          "type": {"name": "string", "args": {"length": 20}},
          "options": {"nullable": false}
        },
        {
          "name": "total",
          "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}},
          "options": {"nullable": false}
        }
      ],
      "indexes": [
        {"name": "idx_customer", "columns": ["customer_id"], "unique": false}
      ]
    },
    {
      "name": "order_items",
      "columns": [
        {
          "name": "id",
          "type": "integer",
          "options": {"primary_key": true, "autoincrement": true}
        },
        {
          "name": "order_id",
          "type": "integer",
          "options": {
            "nullable": false,
            "foreign_key": {"table": "orders", "column": "id", "ondelete": "CASCADE"}
          }
        },
        {
          "name": "product_id",
          "type": "integer",
          "options": {
            "nullable": false,
            "foreign_key": {"table": "products", "column": "id"}
          }
        },
        {
          "name": "quantity",
          "type": "integer",
          "options": {"nullable": false}
        },
        {
          "name": "price",
          "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}},
          "options": {"nullable": false}
        }
      ],
      "indexes": [
        {"name": "idx_order", "columns": ["order_id"], "unique": false},
        {"name": "idx_product", "columns": ["product_id"], "unique": false}
      ]
    }
  ],
  "populate": [
    {
      "name": "customers",
      "count": 500,
      "fields": [
        {"name": "first_name", "generator": "first_name"},
        {"name": "last_name", "generator": "last_name"},
        {"name": "email", "generator": "email"},
        {"name": "phone", "generator": "phone_number"},
        {"name": "address", "generator": "street_address"},
        {"name": "city", "generator": "city"},
        {"name": "state", "generator": "state"},
        {"name": "postcode", "generator": "postcode"},
        {"name": "created_at", "generator": "past_date"}
      ]
    },
    {
      "name": "categories",
      "count": 10,
      "fields": [
        {"name": "name", "generator": "word"},
        {"name": "description", "generator": "sentence"}
      ]
    },
    {
      "name": "products",
      "count": 200,
      "fields": [
        {"name": "category_id", "generator": "foreign_key", "args": {"table": "categories"}},
        {"name": "name", "generator": "sentence"},
        {"name": "description", "generator": "paragraph"},
        {"name": "price", "generator": "random_float", "args": {"min": 9.99, "max": 999.99, "decimals": 2}},
        {"name": "stock", "generator": "random_int", "args": {"min": 0, "max": 500}},
        {"name": "sku", "generator": "uuid"},
        {"name": "created_at", "generator": "past_date"}
      ]
    },
    {
      "name": "orders",
      "count": 1000,
      "fields": [
        {"name": "customer_id", "generator": "foreign_key", "args": {"table": "customers"}},
        {"name": "order_date", "generator": "past_date"},
        {"name": "status", "generator": "random_from", "args": {"choices": ["pending", "processing", "shipped", "delivered", "cancelled"]}},
        {"name": "total", "generator": "random_float", "args": {"min": 10.0, "max": 5000.0, "decimals": 2}}
      ]
    },
    {
      "name": "order_items",
      "count": 3000,
      "fields": [
        {"name": "order_id", "generator": "foreign_key", "args": {"table": "orders"}},
        {"name": "product_id", "generator": "foreign_key", "args": {"table": "products"}},
        {"name": "quantity", "generator": "random_int", "args": {"min": 1, "max": 10}},
        {"name": "price", "generator": "random_float", "args": {"min": 9.99, "max": 999.99, "decimals": 2}}
      ]
    }
  ]
}
```

**Usage:**

```python
from fakestack import Fakestack

faker = Fakestack.from_file('ecommerce.json')
faker.run()
```

---

## Blog Platform

Schema for a blogging platform with users, posts, comments, and tags.

```json
{
  "database": {
    "dbtype": "mysql",
    "drivername": "mysql+mysqlconnector",
    "username": "root",
    "password": "password",
    "host": "localhost",
    "port": 3306,
    "database": "blog"
  },
  "tables": [
    {
      "name": "users",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "username", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}},
        {"name": "email", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false, "unique": true}},
        {"name": "password", "type": {"name": "string", "args": {"length": 255}}, "options": {"nullable": false}},
        {"name": "bio", "type": "text", "options": {}},
        {"name": "avatar_url", "type": {"name": "string", "args": {"length": 255}}, "options": {}},
        {"name": "joined_at", "type": "datetime", "options": {"nullable": false}}
      ],
      "indexes": [
        {"name": "idx_username", "columns": ["username"], "unique": true}
      ]
    },
    {
      "name": "posts",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "user_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "users", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "title", "type": {"name": "string", "args": {"length": 200}}, "options": {"nullable": false}},
        {"name": "slug", "type": {"name": "string", "args": {"length": 200}}, "options": {"nullable": false, "unique": true}},
        {"name": "content", "type": "text", "options": {"nullable": false}},
        {"name": "excerpt", "type": "text", "options": {}},
        {"name": "published", "type": "boolean", "options": {"nullable": false}},
        {"name": "views", "type": "integer", "options": {"nullable": false}},
        {"name": "created_at", "type": "datetime", "options": {"nullable": false}},
        {"name": "updated_at", "type": "datetime", "options": {}}
      ],
      "indexes": [
        {"name": "idx_user", "columns": ["user_id"], "unique": false},
        {"name": "idx_slug", "columns": ["slug"], "unique": true}
      ]
    },
    {
      "name": "comments",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "post_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "posts", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "user_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "users", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "content", "type": "text", "options": {"nullable": false}},
        {"name": "created_at", "type": "datetime", "options": {"nullable": false}}
      ],
      "indexes": [
        {"name": "idx_post", "columns": ["post_id"], "unique": false},
        {"name": "idx_user", "columns": ["user_id"], "unique": false}
      ]
    },
    {
      "name": "tags",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "name", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}},
        {"name": "slug", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}}
      ],
      "indexes": []
    },
    {
      "name": "post_tags",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "post_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "posts", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "tag_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "tags", "column": "id", "ondelete": "CASCADE"}}}
      ],
      "indexes": [
        {"name": "idx_post_tag", "columns": ["post_id", "tag_id"], "unique": true}
      ]
    }
  ],
  "populate": [
    {
      "name": "users",
      "count": 100,
      "fields": [
        {"name": "username", "generator": "user_name"},
        {"name": "email", "generator": "email"},
        {"name": "password", "generator": "password"},
        {"name": "bio", "generator": "paragraph"},
        {"name": "avatar_url", "generator": "url"},
        {"name": "joined_at", "generator": "past_date"}
      ]
    },
    {
      "name": "posts",
      "count": 500,
      "fields": [
        {"name": "user_id", "generator": "foreign_key", "args": {"table": "users"}},
        {"name": "title", "generator": "sentence"},
        {"name": "slug", "generator": "slug"},
        {"name": "content", "generator": "text", "args": {"max_chars": 2000}},
        {"name": "excerpt", "generator": "paragraph"},
        {"name": "published", "generator": "boolean"},
        {"name": "views", "generator": "random_int", "args": {"min": 0, "max": 10000}},
        {"name": "created_at", "generator": "past_date"},
        {"name": "updated_at", "generator": "past_date"}
      ]
    },
    {
      "name": "comments",
      "count": 2000,
      "fields": [
        {"name": "post_id", "generator": "foreign_key", "args": {"table": "posts"}},
        {"name": "user_id", "generator": "foreign_key", "args": {"table": "users"}},
        {"name": "content", "generator": "paragraph"},
        {"name": "created_at", "generator": "past_date"}
      ]
    },
    {
      "name": "tags",
      "count": 50,
      "fields": [
        {"name": "name", "generator": "word"},
        {"name": "slug", "generator": "slug"}
      ]
    },
    {
      "name": "post_tags",
      "count": 1500,
      "fields": [
        {"name": "post_id", "generator": "foreign_key", "args": {"table": "posts"}},
        {"name": "tag_id", "generator": "foreign_key", "args": {"table": "tags"}}
      ]
    }
  ]
}
```

---

## Social Network

Simplified social network with users, friendships, and posts.

```json
{
  "database": {
    "dbtype": "postgresql",
    "drivername": "postgresql+psycopg2",
    "username": "postgres",
    "password": "postgres",
    "host": "localhost",
    "port": 5432,
    "database": "social"
  },
  "tables": [
    {
      "name": "users",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "username", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}},
        {"name": "email", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false, "unique": true}},
        {"name": "full_name", "type": {"name": "string", "args": {"length": 100}}, "options": {}},
        {"name": "bio", "type": "text", "options": {}},
        {"name": "profile_picture", "type": {"name": "string", "args": {"length": 255}}, "options": {}},
        {"name": "created_at", "type": "datetime", "options": {"nullable": false}}
      ],
      "indexes": []
    },
    {
      "name": "friendships",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "user_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "users", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "friend_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "users", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "status", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}},
        {"name": "created_at", "type": "datetime", "options": {"nullable": false}}
      ],
      "indexes": [
        {"name": "idx_user_friend", "columns": ["user_id", "friend_id"], "unique": true}
      ]
    },
    {
      "name": "posts",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "user_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "users", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "content", "type": "text", "options": {"nullable": false}},
        {"name": "likes", "type": "integer", "options": {"nullable": false}},
        {"name": "created_at", "type": "datetime", "options": {"nullable": false}}
      ],
      "indexes": [
        {"name": "idx_user", "columns": ["user_id"], "unique": false}
      ]
    }
  ],
  "populate": [
    {
      "name": "users",
      "count": 1000,
      "fields": [
        {"name": "username", "generator": "user_name"},
        {"name": "email", "generator": "email"},
        {"name": "full_name", "generator": "name"},
        {"name": "bio", "generator": "paragraph"},
        {"name": "profile_picture", "generator": "url"},
        {"name": "created_at", "generator": "past_date"}
      ]
    },
    {
      "name": "friendships",
      "count": 5000,
      "fields": [
        {"name": "user_id", "generator": "foreign_key", "args": {"table": "users"}},
        {"name": "friend_id", "generator": "foreign_key", "args": {"table": "users"}},
        {"name": "status", "generator": "random_from", "args": {"choices": ["pending", "accepted", "rejected"]}},
        {"name": "created_at", "generator": "past_date"}
      ]
    },
    {
      "name": "posts",
      "count": 10000,
      "fields": [
        {"name": "user_id", "generator": "foreign_key", "args": {"table": "users"}},
        {"name": "content", "generator": "paragraph"},
        {"name": "likes", "generator": "random_int", "args": {"min": 0, "max": 500}},
        {"name": "created_at", "generator": "past_date"}
      ]
    }
  ]
}
```

---

## Task Management

Project and task tracking system.

```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "tasks.db"
  },
  "tables": [
    {
      "name": "projects",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "name", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false}},
        {"name": "description", "type": "text", "options": {}},
        {"name": "start_date", "type": "date", "options": {}},
        {"name": "end_date", "type": "date", "options": {}},
        {"name": "status", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}}
      ],
      "indexes": []
    },
    {
      "name": "tasks",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "project_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "projects", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "title", "type": {"name": "string", "args": {"length": 200}}, "options": {"nullable": false}},
        {"name": "description", "type": "text", "options": {}},
        {"name": "priority", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}},
        {"name": "status", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}},
        {"name": "due_date", "type": "date", "options": {}},
        {"name": "created_at", "type": "datetime", "options": {"nullable": false}}
      ],
      "indexes": [
        {"name": "idx_project", "columns": ["project_id"], "unique": false}
      ]
    }
  ],
  "populate": [
    {
      "name": "projects",
      "count": 50,
      "fields": [
        {"name": "name", "generator": "sentence"},
        {"name": "description", "generator": "paragraph"},
        {"name": "start_date", "generator": "past_date"},
        {"name": "end_date", "generator": "future_date"},
        {"name": "status", "generator": "random_from", "args": {"choices": ["planning", "active", "on_hold", "completed"]}}
      ]
    },
    {
      "name": "tasks",
      "count": 500,
      "fields": [
        {"name": "project_id", "generator": "foreign_key", "args": {"table": "projects"}},
        {"name": "title", "generator": "sentence"},
        {"name": "description", "generator": "paragraph"},
        {"name": "priority", "generator": "random_from", "args": {"choices": ["low", "medium", "high", "urgent"]}},
        {"name": "status", "generator": "random_from", "args": {"choices": ["todo", "in_progress", "review", "done"]}},
        {"name": "due_date", "generator": "future_date"},
        {"name": "created_at", "generator": "past_date"}
      ]
    }
  ]
}
```

---

## Event Registration

Event management with registrations and tickets.

```json
{
  "database": {
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "events.db"
  },
  "tables": [
    {
      "name": "events",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "name", "type": {"name": "string", "args": {"length": 200}}, "options": {"nullable": false}},
        {"name": "description", "type": "text", "options": {}},
        {"name": "location", "type": {"name": "string", "args": {"length": 200}}, "options": {}},
        {"name": "event_date", "type": "datetime", "options": {"nullable": false}},
        {"name": "capacity", "type": "integer", "options": {"nullable": false}},
        {"name": "price", "type": {"name": "decimal", "args": {"precision": 10, "scale": 2}}, "options": {"nullable": false}}
      ],
      "indexes": []
    },
    {
      "name": "attendees",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "first_name", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false}},
        {"name": "last_name", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false}},
        {"name": "email", "type": {"name": "string", "args": {"length": 100}}, "options": {"nullable": false, "unique": true}},
        {"name": "phone", "type": {"name": "string", "args": {"length": 20}}, "options": {}}
      ],
      "indexes": []
    },
    {
      "name": "registrations",
      "columns": [
        {"name": "id", "type": "integer", "options": {"primary_key": true, "autoincrement": true}},
        {"name": "event_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "events", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "attendee_id", "type": "integer", "options": {"nullable": false, "foreign_key": {"table": "attendees", "column": "id", "ondelete": "CASCADE"}}},
        {"name": "ticket_number", "type": {"name": "string", "args": {"length": 50}}, "options": {"nullable": false, "unique": true}},
        {"name": "registered_at", "type": "datetime", "options": {"nullable": false}},
        {"name": "status", "type": {"name": "string", "args": {"length": 20}}, "options": {"nullable": false}}
      ],
      "indexes": [
        {"name": "idx_event_attendee", "columns": ["event_id", "attendee_id"], "unique": true}
      ]
    }
  ],
  "populate": [
    {
      "name": "events",
      "count": 20,
      "fields": [
        {"name": "name", "generator": "sentence"},
        {"name": "description", "generator": "paragraph"},
        {"name": "location", "generator": "address"},
        {"name": "event_date", "generator": "future_date"},
        {"name": "capacity", "generator": "random_int", "args": {"min": 50, "max": 1000}},
        {"name": "price", "generator": "random_float", "args": {"min": 0.0, "max": 500.0, "decimals": 2}}
      ]
    },
    {
      "name": "attendees",
      "count": 500,
      "fields": [
        {"name": "first_name", "generator": "first_name"},
        {"name": "last_name", "generator": "last_name"},
        {"name": "email", "generator": "email"},
        {"name": "phone", "generator": "phone_number"}
      ]
    },
    {
      "name": "registrations",
      "count": 800,
      "fields": [
        {"name": "event_id", "generator": "foreign_key", "args": {"table": "events"}},
        {"name": "attendee_id", "generator": "foreign_key", "args": {"table": "attendees"}},
        {"name": "ticket_number", "generator": "uuid"},
        {"name": "registered_at", "generator": "past_date"},
        {"name": "status", "generator": "random_from", "args": {"choices": ["confirmed", "pending", "cancelled"]}}
      ]
    }
  ]
}
```

---

## See Also

- [Schema Reference](schema-reference.md) - Complete schema documentation
- [Data Generators](generators.md) - All available generators
- [Database Support](databases.md) - Database-specific information
- [Getting Started](getting-started.md) - Quick start guide
