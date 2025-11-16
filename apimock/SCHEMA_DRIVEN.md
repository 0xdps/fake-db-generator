# Schema-Driven API Development

This system allows you to add new API resources by simply creating a JSON Schema file. No code changes required!

## How It Works

```
┌─────────────────────┐
│  Add JSON Schema    │
│  (shared/schemas/)  │
└──────────┬──────────┘
           │
           ├──────────────────────┬──────────────────────┐
           │                      │                      │
           ▼                      ▼                      ▼
    ┌──────────┐          ┌──────────┐          ┌──────────┐
    │   API    │          │   Web    │          │  Types   │
    │ Routes   │          │  Docs    │          │  (.ts)   │
    │   (Go)   │          │  (Auto)  │          │  (Auto)  │
    └──────────┘          └──────────┘          └──────────┘
```

## Adding a New Resource

1. **Create a JSON Schema** in `apimock/shared/schemas/`
2. **Restart the API** (or it will auto-reload in dev mode)
3. **Generate TypeScript types** by running `npm run generate-types`

That's it! Your new resource is now available at `/api/{resource}` and `/api/{resource}/:id`

## Schema Format

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "ResourceName",
  "type": "object",
  "description": "Resource description",
  "x-resource": {
    "name": "resources",
    "singular": "resource",
    "description": "Displayed in API documentation"
  },
  "properties": {
    "id": {
      "type": "integer",
      "description": "Unique identifier",
      "x-generator": "random_int",
      "x-generator-params": {
        "min": 1,
        "max": 10000
      }
    },
    "name": {
      "type": "string",
      "description": "Resource name",
      "x-generator": "word"
    }
  },
  "required": ["id", "name"]
}
```

### Custom Extensions

- **`x-resource`**: Metadata about the resource
  - `name`: Plural name for the API endpoint
  - `singular`: Singular form (for documentation)
  - `description`: Human-readable description
  - `routes`: **NEW!** Custom route configuration
    - `path`: Custom API path (e.g., `/v1/users`)
    - `methods`: Allowed HTTP methods (e.g., `["GET", "POST"]`)
    - `aliases`: Additional paths (e.g., `["/tasks", "/todos"]`)

- **`x-generator`**: Specifies which fake data generator to use
  - See [Supported Generators](#supported-generators) below

- **`x-generator-params`**: Parameters for the generator (e.g., min/max for numbers)

📘 **See [CUSTOM_ROUTES.md](./CUSTOM_ROUTES.md) for detailed route customization guide**

## Supported Generators

### Personal
- `name`, `first_name`, `last_name`
- `user_name`, `username`
- `email`, `password`, `gender`

### Address
- `address`, `street_address`, `city`, `state`, `country`
- `postcode`, `zip_code`
- `latitude`, `longitude`

### Company
- `company`, `job`, `catch_phrase`

### Internet
- `url`, `domain_name`
- `ipv4`, `ipv6`, `mac_address`

### Dates
- `date`, `date_time`
- `past_date`, `future_date`

### Text
- `text`, `sentence`, `paragraph`, `word`

### Numbers
- `random_int` - accepts `min` and `max` params
- `random_digit`, `random_number`

### Other
- `phone_number`, `phone`
- `boolean`
- `uuid`

## Example: Adding a "Todo" Resource

1. Create `apimock/shared/schemas/todo.json`:

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Todo",
  "type": "object",
  "x-resource": {
    "name": "todos",
    "singular": "todo",
    "description": "Todo items"
  },
  "properties": {
    "id": {
      "type": "integer",
      "x-generator": "random_int",
      "x-generator-params": { "min": 1, "max": 10000 }
    },
    "user_id": {
      "type": "integer",
      "x-generator": "random_int",
      "x-generator-params": { "min": 1, "max": 100 }
    },
    "title": {
      "type": "string",
      "x-generator": "sentence"
    },
    "completed": {
      "type": "boolean",
      "x-generator": "boolean"
    }
  },
  "required": ["id", "title"]
}
```

2. Restart the server:
```bash
cd apimock/api && go run main.go
```

3. Generate types:
```bash
npm run generate-types
```

4. Test the new endpoint:
```bash
curl http://localhost:8080/api/todos?count=5
```

That's it! You now have:
- ✅ `GET /api/todos` - List all todos
- ✅ `GET /api/todos/:id` - Get single todo
- ✅ `GET /api/todos/meta` - Resource metadata
- ✅ TypeScript types in `web/types/api.ts`

## Available Endpoints

For each resource, you automatically get:

- `GET /api/{resource}?count=N` - Get N items (default 10, max 1000)
- `GET /api/{resource}/:id` - Get single item by ID
- `GET /api/{resource}/meta` - Get resource metadata

## Scripts

- `npm run generate-types` - Generate TypeScript types from schemas
- `npm run dev` - Start development server (auto-reloads on schema changes)

## Current Resources

The API currently provides these resources:

- **Users** (`/api/users`) - User accounts
- **Posts** (`/api/posts`) - Blog posts
- **Products** (`/api/products`) - E-commerce products
- **Comments** (`/api/comments`) - Comments on posts

Check http://localhost:8080/ to see all available resources.
