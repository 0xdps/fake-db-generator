# Custom Routes in Schemas

## Overview

You can now define custom routes directly in your JSON Schema files! This allows you to:

- **Custom paths** - Use versioned paths like `/v1/users` or `/v2/posts`
- **Route aliases** - Create multiple paths for the same resource
- **HTTP methods** - Specify which methods are supported (GET, POST, PUT, DELETE)

## Schema Configuration

Add the `routes` object to `x-resource` in your schema:

```json
{
  "x-resource": {
    "name": "todos",
    "singular": "todo",
    "description": "Todo items",
    "routes": {
      "path": "/v1/todos",
      "methods": ["GET", "POST"],
      "aliases": ["/tasks", "/todo-items"]
    }
  }
}
```

### Route Properties

| Property | Type | Description | Example |
|----------|------|-------------|---------|
| `path` | string | Custom API path | `"/v1/users"`, `"/blog/posts"` |
| `methods` | array | Allowed HTTP methods | `["GET", "POST", "DELETE"]` |
| `aliases` | array | Additional paths for this resource | `["/tasks", "/todos"]` |

All properties are **optional**. If not specified:
- `path` defaults to `/{resource-name}`
- `methods` defaults to `["GET"]`
- `aliases` defaults to none

## Examples

### Example 1: Versioned API

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Todo",
  "x-resource": {
    "name": "todos",
    "routes": {
      "path": "/v1/todos"
    }
  },
  "properties": { ... }
}
```

**Result:**
- ✅ `GET /api/v1/todos`
- ✅ `GET /api/v1/todos/:id`
- ✅ `GET /api/v1/todos/meta`

### Example 2: Multiple Aliases

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Todo",
  "x-resource": {
    "name": "todos",
    "routes": {
      "path": "/v1/todos",
      "aliases": ["/tasks", "/todo-items", "/checklist"]
    }
  },
  "properties": { ... }
}
```

**Result:**
- ✅ `GET /api/v1/todos` (main path)
- ✅ `GET /api/tasks` (alias 1)
- ✅ `GET /api/todo-items` (alias 2)
- ✅ `GET /api/checklist` (alias 3)
- ... and `/:id` + `/meta` for each

### Example 3: Method Support

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Review",
  "x-resource": {
    "name": "reviews",
    "routes": {
      "path": "/reviews",
      "methods": ["GET", "POST", "PUT", "DELETE"]
    }
  },
  "properties": { ... }
}
```

**Result:**
- Indicates support for all CRUD operations
- Currently all generate the same fake data (future: method-specific behavior)

## Real-World Use Cases

### Use Case 1: API Versioning

Keep old and new versions running simultaneously:

```json
// v1-users.json
{
  "x-resource": {
    "name": "users_v1",
    "routes": { "path": "/v1/users" }
  }
}

// v2-users.json
{
  "x-resource": {
    "name": "users_v2",
    "routes": { "path": "/v2/users" }
  }
}
```

### Use Case 2: Legacy Compatibility

Support old URL patterns while transitioning:

```json
{
  "x-resource": {
    "name": "articles",
    "routes": {
      "path": "/blog/articles",
      "aliases": ["/posts", "/news", "/articles"]
    }
  }
}
```

Clients can use any of these:
- `/api/blog/articles` (new)
- `/api/posts` (legacy)
- `/api/news` (legacy)
- `/api/articles` (short form)

### Use Case 3: Semantic URLs

Create more intuitive paths:

```json
// Instead of /api/reviews
{
  "x-resource": {
    "name": "reviews",
    "routes": {
      "path": "/product-reviews",
      "aliases": ["/ratings", "/feedback"]
    }
  }
}
```

## Live Demo

We tested this with the `todo` schema:

**Schema:**
```json
{
  "x-resource": {
    "name": "todos",
    "routes": {
      "path": "/v1/todos",
      "aliases": ["/tasks", "/todo-items"]
    }
  }
}
```

**Tests:**
```bash
# Main path
curl http://localhost:8080/api/v1/todos?count=1
# ✅ Works! Returns todo data

# Alias 1
curl http://localhost:8080/api/tasks?count=1
# ✅ Works! Same data, different path

# Alias 2
curl http://localhost:8080/api/todo-items?count=1
# ✅ Works! Same data, another path
```

## Important Notes

### Route Ordering

Routes are registered in this order:
1. **Nested routes first** (routes containing `:` parameters)
2. **Regular routes second**

This prevents Gin router conflicts.

### Nested Routes Limitation

Due to Gin's router limitations, you **cannot** have both:
- `/products/:id`
- `/products/:productId/reviews`

**Workaround:** Use separate top-level routes:
```json
// reviews.json
{
  "x-resource": {
    "name": "reviews",
    "routes": {
      "path": "/reviews"  // Not nested
    }
  }
}
```

Then filter by query parameter: `/api/reviews?product_id=123`

## Auto-Generated Endpoints

For each schema with custom routes, you automatically get:

| Endpoint | Purpose | Example |
|----------|---------|---------|
| `{path}` | Collection | `/api/v1/todos?count=10` |
| `{path}/:id` | Single item | `/api/v1/todos/42` |
| `{path}/meta` | Metadata | `/api/v1/todos/meta` |

Plus the same for each alias!

## Migration from Default Routes

### Before (no custom routes)
```json
{
  "x-resource": {
    "name": "todos"
  }
}
```
Generates: `/api/todos`

### After (with custom routes)
```json
{
  "x-resource": {
    "name": "todos",
    "routes": {
      "path": "/v1/todos",
      "aliases": ["/todos"]  // Keep backward compatibility
    }
  }
}
```
Generates: `/api/v1/todos` + `/api/todos` (both work!)

## Server Logs

The server logs show registered routes on startup:

```
2025/11/16 15:20:50 Loaded 6 schemas
2025/11/16 15:20:50 Registered routes: todos -> /v1/todos
2025/11/16 15:20:50   + alias: /tasks
2025/11/16 15:20:50   + alias: /todo-items
2025/11/16 15:20:50 Registered routes: reviews -> /reviews
2025/11/16 15:20:50   + alias: /product-reviews
2025/11/16 15:20:50   + alias: /v2/reviews
```

## Summary

**With custom routes, you can:**

✅ Version your API (`/v1/`, `/v2/`)  
✅ Support legacy URLs (aliases)  
✅ Create semantic paths (`/product-reviews` instead of `/reviews`)  
✅ Keep backward compatibility (main path + aliases)  
✅ Define supported HTTP methods  

**All without writing a single line of code - just edit the schema!** 🚀
