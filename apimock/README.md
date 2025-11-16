# apimock.codes

Free mock API service for testing and development - An alternative to JSONPlaceholder and httpbin.

## Project Structure

```
apimock/
├── api/          # Go API server (api.apimock.codes)
├── web/          # Next.js website (apimock.codes)
└── shared/       # Shared schemas and types
```

## Features

✅ **RESTful Mock APIs**
- `/api/users` - User data
- `/api/posts` - Blog posts
- `/api/products` - Product catalog
- More coming soon...

✅ **Powered by Fakestack**
- Realistic fake data generation
- Consistent data across requests
- Multiple data generators

✅ **Developer Friendly**
- CORS enabled
- No authentication required
- Rate limiting friendly
- Query parameters (`?count=N`)

## Quick Start

### API Development

```bash
cd api
go run main.go
```

Server: `http://localhost:8080`

### Web Development

```bash
cd web
npm install
npm run dev
```

Website: `http://localhost:3000`

### Full Stack Development

Start both services:

```bash
# Terminal 1 - API
cd api && go run main.go

# Terminal 2 - Website
cd web && npm run dev
```

Then visit:
- Website: http://localhost:3000
- API: http://localhost:8080
- Documentation: http://localhost:3000/docs
- Playground: http://localhost:3000/playground

## API Examples

```bash
# Get 10 users
curl http://localhost:8080/api/users?count=10

# Get single user
curl http://localhost:8080/api/users/123

# Get products
curl http://localhost:8080/api/products?count=50

# Get posts
curl http://localhost:8080/api/posts
```

## Deployment

### Quick Deploy to Vercel

```bash
# Automated deployment (recommended)
cd apimock
./deploy.sh

# Or manually deploy each service
cd api && vercel --prod
cd web && vercel --prod
```

See [VERCEL_SETUP.md](VERCEL_SETUP.md) for detailed deployment guide.

### API (Vercel)
- Project root: `api/`
- Framework: Other (Go)
- Domain: `api.apimock.codes`

### Website (Vercel)
- Project root: `web/`
- Framework: Next.js
- Domain: `apimock.codes`

## Technology Stack

- **API**: Go + Gin + Fakestack Generator
- **Website**: Next.js + React + Tailwind CSS
- **Hosting**: Vercel
- **Data**: Fakestack (no database needed)

## Related Projects

- [fakestack](../README.md) - Database population tool (Python, Node.js, Go, Homebrew)
- [fakestack docs](https://fakestack.readthedocs.io) - Package documentation

## License

MIT
