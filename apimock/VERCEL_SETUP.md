# Vercel Deployment Guide - apimock.codes

This guide walks you through deploying both the API and website to Vercel.

## Project Architecture

```
apimock/
├── api/          # Go API → api.apimock.codes
└── web/          # Next.js → apimock.codes
```

## Prerequisites

1. **Vercel Account**: Sign up at https://vercel.com
2. **Vercel CLI** (optional but recommended):
   ```bash
   npm install -g vercel
   ```
3. **GitHub Repository**: Your repo at https://github.com/0xdps/fake-stack

## Deployment Steps

### Step 1: Deploy the API First

The website depends on the API, so deploy the API first.

#### Option A: Using Vercel Dashboard (Recommended)

1. Go to https://vercel.com/new
2. Import your GitHub repository: `0xdps/fake-stack`
3. Configure the API project:
   - **Project Name**: `apimock-api` (or any name)
   - **Framework Preset**: Other
   - **Root Directory**: `apimock/api`
   - **Build Command**: Leave empty (Go doesn't need build)
   - **Output Directory**: Leave empty
   - **Install Command**: Leave empty

4. **Environment Variables**: None needed for API

5. Click **Deploy**

6. After deployment, note your API URL (e.g., `apimock-api.vercel.app`)

7. **Add Custom Domain**:
   - Go to Project Settings → Domains
   - Add domain: `api.apimock.codes`
   - Configure DNS:
     ```
     Type: CNAME
     Name: api
     Value: cname.vercel-dns.com
     ```

#### Option B: Using Vercel CLI

```bash
# Navigate to API directory
cd apimock/api

# Deploy
vercel --prod

# Add custom domain
vercel domains add api.apimock.codes
```

### Step 2: Deploy the Website

After the API is deployed and you have the production URL:

#### Option A: Using Vercel Dashboard

1. Go to https://vercel.com/new
2. Import the same repository: `0xdps/fake-stack`
3. Configure the web project:
   - **Project Name**: `apimock-web` (or any name)
   - **Framework Preset**: Next.js
   - **Root Directory**: `apimock/web`
   - **Build Command**: `npm run build` (auto-detected)
   - **Output Directory**: `.next` (auto-detected)
   - **Install Command**: `npm install` (auto-detected)

4. **Environment Variables**:
   ```
   Key: NEXT_PUBLIC_API_URL
   Value: https://api.apimock.codes
   ```

5. Click **Deploy**

6. After deployment, note your web URL (e.g., `apimock-web.vercel.app`)

7. **Add Custom Domain**:
   - Go to Project Settings → Domains
   - Add domain: `apimock.codes`
   - Configure DNS:
     ```
     Type: CNAME
     Name: @  (or www)
     Value: cname.vercel-dns.com
     ```

#### Option B: Using Vercel CLI

```bash
# Navigate to web directory
cd apimock/web

# Set environment variable
vercel env add NEXT_PUBLIC_API_URL production

# When prompted, enter: https://api.apimock.codes

# Deploy
vercel --prod

# Add custom domain
vercel domains add apimock.codes
```

## Vercel Configuration Files

### API: `apimock/api/vercel.json`

```json
{
  "version": 2,
  "builds": [
    {
      "src": "main.go",
      "use": "@vercel/go"
    }
  ],
  "routes": [
    {
      "src": "/(.*)",
      "dest": "main.go"
    }
  ]
}
```

### Web: Already configured via Next.js

Next.js projects are automatically configured by Vercel. The `next.config.js` handles everything.

## DNS Configuration

### For `apimock.codes` (root domain)

**Option 1: Use Vercel Nameservers (Recommended)**
- Update your domain registrar to use Vercel nameservers
- Vercel will manage all DNS records

**Option 2: Use CNAME records**

At your domain registrar (e.g., GoDaddy, Namecheap):

```
# Root domain
Type: A
Name: @
Value: 76.76.21.21

# API subdomain
Type: CNAME
Name: api
Value: cname.vercel-dns.com

# WWW subdomain (optional)
Type: CNAME
Name: www
Value: cname.vercel-dns.com
```

## Post-Deployment Checklist

### API Deployment
- [ ] API is accessible at https://api.apimock.codes
- [ ] Test endpoint: `curl https://api.apimock.codes/`
- [ ] CORS is working (test from browser)
- [ ] All resources load: `/api/users`, `/api/posts`, etc.
- [ ] Schemas are accessible: `/api/users/meta`

### Website Deployment
- [ ] Website is accessible at https://apimock.codes
- [ ] Landing page loads correctly
- [ ] Documentation page shows all resources
- [ ] Playground connects to API successfully
- [ ] Test making API calls from playground
- [ ] All links work (header, footer)
- [ ] Mobile responsiveness verified

## Testing After Deployment

```bash
# Test API
curl https://api.apimock.codes/
curl https://api.apimock.codes/api/users?count=5

# Test Website
open https://apimock.codes
open https://apimock.codes/docs
open https://apimock.codes/playground
```

## Environment Variables Reference

### API (api.apimock.codes)
- No environment variables needed

### Website (apimock.codes)
- `NEXT_PUBLIC_API_URL`: Your API domain (e.g., `https://api.apimock.codes`)

## Updating Environment Variables

### Via Vercel Dashboard
1. Go to your project
2. Settings → Environment Variables
3. Edit `NEXT_PUBLIC_API_URL`
4. Trigger a redeployment

### Via Vercel CLI
```bash
# List variables
vercel env ls

# Add/update variable
vercel env add NEXT_PUBLIC_API_URL production

# Pull variables locally
vercel env pull
```

## Continuous Deployment

Once linked to GitHub, Vercel automatically:
- ✅ Deploys on every push to `trunk` (or `main`)
- ✅ Creates preview deployments for PRs
- ✅ Runs builds and shows logs
- ✅ Updates custom domains

### Trigger Manual Deployment

```bash
# Redeploy API
cd apimock/api
vercel --prod

# Redeploy website
cd apimock/web
vercel --prod
```

## Troubleshooting

### API Issues

**Problem**: API returns 404
- Check `vercel.json` is in `apimock/api/` directory
- Verify root directory is set to `apimock/api`
- Check build logs in Vercel dashboard

**Problem**: CORS errors
- Ensure `main.go` has CORS middleware enabled
- Check API is responding to OPTIONS requests

### Website Issues

**Problem**: Website can't connect to API
- Verify `NEXT_PUBLIC_API_URL` environment variable is set
- Check API is deployed and accessible
- Look at browser console for errors

**Problem**: 404 on routes
- Verify root directory is `apimock/web`
- Check Next.js build succeeded
- Review build logs

### General Issues

**Problem**: Build fails
```bash
# Check locally first
cd apimock/api && go run main.go
cd apimock/web && npm run build
```

**Problem**: Environment variables not updating
- Redeploy after changing env vars
- Clear Vercel cache if needed

## Custom Domain Verification

After adding custom domains, verify with:

```bash
# Check DNS propagation
dig api.apimock.codes
dig apimock.codes

# Test HTTPS
curl -I https://api.apimock.codes
curl -I https://apimock.codes
```

## Project URLs

- **API Production**: https://api.apimock.codes
- **Website Production**: https://apimock.codes
- **Documentation**: https://apimock.codes/docs
- **Playground**: https://apimock.codes/playground
- **GitHub**: https://github.com/0xdps/fake-stack

## Quick Deploy Commands

```bash
# Deploy both services
cd /path/to/fake-stack/apimock

# Deploy API
(cd api && vercel --prod)

# Deploy Website  
(cd web && vercel --prod)
```

## Support

- **Vercel Docs**: https://vercel.com/docs
- **Vercel Go Runtime**: https://vercel.com/docs/runtimes#official-runtimes/go
- **Next.js on Vercel**: https://vercel.com/docs/frameworks/nextjs

---

**Ready to Deploy?** Follow Step 1 and Step 2 above! 🚀
