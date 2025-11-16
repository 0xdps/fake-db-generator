# apimock.codes Website

Modern Next.js website for the apimock.codes schema-driven mock API service.

## Features

- 🎨 **Modern UI** - Beautiful gradient design with Tailwind CSS
- 📱 **Fully Responsive** - Works on all devices
- 🎮 **Interactive Playground** - Test API endpoints in-browser
- 📚 **Auto-Generated Docs** - Documentation from JSON schemas
- ⚡️ **Fast** - Built with Next.js 14 and App Router
- 🌙 **Dark Mode** - Optimized for dark theme

## Getting Started

### Install Dependencies

```bash
npm install
```

### Development Server

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) to see the website.

### Build for Production

```bash
npm run build
npm start
```

## Environment Variables

Create a `.env.local` file:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

For production:

```env
NEXT_PUBLIC_API_URL=https://your-api-domain.com
```

## Project Structure

```
web/
├── app/
│   ├── page.tsx          # Landing page
│   ├── docs/page.tsx     # Documentation page
│   ├── playground/page.tsx # API playground
│   ├── layout.tsx        # Root layout
│   └── globals.css       # Global styles
├── components/
│   ├── Header.tsx        # Navigation header
│   ├── Footer.tsx        # Footer
│   ├── ResourceCard.tsx  # Resource cards
│   └── CodeExample.tsx   # Code snippet component
└── public/               # Static assets
```

## Pages

### Landing Page (`/`)
- Hero section with CTA buttons
- Feature highlights (6 key features)
- Quick start example
- Available resources showcase
- Use cases section

### Documentation (`/docs`)
- Quick start guide
- Code examples in multiple languages
- Endpoint reference
- Complete resource schemas
- Feature documentation

### Playground (`/playground`)
- Interactive API testing
- Resource selection
- Endpoint type switcher (collection/single/meta)
- Live response viewer
- Auto-generated code examples
- Copy to clipboard

## Deployment

### Vercel (Recommended)

```bash
vercel
```

Or connect your GitHub repo to Vercel for automatic deployments.

### Environment Variables for Production

Set in Vercel dashboard:
- `NEXT_PUBLIC_API_URL` - Your API domain

## Tech Stack

- **Framework**: Next.js 14 (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **Deployment**: Vercel

## Features Showcase

### Dynamic Resource Loading
All resources are loaded from the API at build time, so adding a new schema automatically updates the website.

### Schema-Driven Documentation
Documentation is generated from JSON schemas in `../shared/schemas/`.

### Interactive Playground
Test all endpoints without writing code:
- Select resource
- Choose endpoint type
- Adjust parameters
- See real-time responses
- Copy code examples

## License

MIT
