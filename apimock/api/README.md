# apimock.codes API

Mock API service for testing and development.

## Development

```bash
go run main.go
```

Server runs on `http://localhost:8080`

## Available Endpoints

### Users
- `GET /api/users` - List users
- `GET /api/users/:id` - Get single user
- Query params: `?count=N` (max 1000)

### Posts
- `GET /api/posts` - List posts
- `GET /api/posts/:id` - Get single post
- Query params: `?count=N` (max 1000)

### Products
- `GET /api/products` - List products
- `GET /api/products/:id` - Get single product
- Query params: `?count=N` (max 1000)

## Examples

```bash
# Get 10 users
curl http://localhost:8080/api/users?count=10

# Get single user
curl http://localhost:8080/api/users/123

# Get 50 products
curl http://localhost:8080/api/products?count=50
```

## Deployment

Deploy to Vercel:

```bash
vercel --prod
```

Configure:
- Root Directory: `api/`
- Framework: Other (Go)
- Domain: `api.apimock.codes`
