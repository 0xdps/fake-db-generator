# Quick Start: Deploy to Vercel NOW

Follow these exact commands to deploy your project to Vercel in 5 minutes.

## Step 1: Login to Vercel

```bash
vercel login
```

This will open your browser. Login with GitHub, GitLab, or email.

## Step 2: Deploy API

```bash
cd /Users/devendrapratapsingh/personal/fake-stack/apimock/api
vercel
```

**What to answer:**
- `Set up and deploy?` → **Y** (yes)
- `Which scope?` → Select your account
- `Link to existing project?` → **N** (no)
- `What's your project's name?` → **apimock-api** (or any name)
- `In which directory is your code located?` → **./** (press Enter)

Vercel will deploy! Note the preview URL (e.g., `https://apimock-api-xxx.vercel.app`)

**Deploy to production:**
```bash
vercel --prod
```

Note your production URL. You'll need it for the website.

## Step 3: Deploy Website

```bash
cd /Users/devendrapratapsingh/personal/fake-stack/apimock/web
```

**Set the API URL environment variable:**
```bash
vercel env add NEXT_PUBLIC_API_URL production
```

When prompted, enter your API production URL from Step 2 (e.g., `https://apimock-api.vercel.app`)

**Deploy:**
```bash
vercel
```

**What to answer:**
- `Set up and deploy?` → **Y** (yes)
- `Which scope?` → Select your account
- `Link to existing project?` → **N** (no)
- `What's your project's name?` → **apimock-web** (or any name)
- `In which directory is your code located?` → **./** (press Enter)

Vercel will build and deploy!

**Deploy to production:**
```bash
vercel --prod
```

## Step 4: Test Your Deployments

```bash
# Test API
curl https://YOUR-API-URL.vercel.app/

# Open website in browser
open https://YOUR-WEB-URL.vercel.app
```

## Step 5: Add Custom Domains (Optional)

In the Vercel dashboard:

1. **For API Project:**
   - Go to Settings → Domains
   - Add: `api.apimock.codes`
   - Configure DNS as shown

2. **For Website Project:**
   - Go to Settings → Domains
   - Add: `apimock.codes`
   - Configure DNS as shown

## One-Command Deployment (Alternative)

After initial setup, use the deployment script:

```bash
cd /Users/devendrapratapsingh/personal/fake-stack/apimock
./deploy.sh
```

This will:
- Deploy API to production
- Extract API URL
- Set environment variable for website
- Deploy website to production
- Show summary with all URLs

## Troubleshooting

### "No framework detected"
This is normal for the API (Go). Just continue.

### "Build failed"
```bash
# Test locally first
cd api && go run main.go
cd web && npm run build
```

### "Environment variable not set"
```bash
cd web
vercel env add NEXT_PUBLIC_API_URL production
# Then redeploy
vercel --prod
```

## What's Deployed?

After successful deployment:

✅ **API** (`https://YOUR-API.vercel.app`)
- Root endpoint: `/`
- Users: `/api/users`
- All 6 resources working

✅ **Website** (`https://YOUR-WEB.vercel.app`)
- Landing page: `/`
- Documentation: `/docs`
- Playground: `/playground`

## Next Steps

1. ✅ Test all endpoints
2. ✅ Try the playground
3. ✅ Check documentation page
4. ⏳ Add custom domains
5. ⏳ Share with the world!

---

**Need help?** See [VERCEL_SETUP.md](VERCEL_SETUP.md) for detailed guide.
