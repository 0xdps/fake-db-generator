# Troubleshooting

Common issues and solutions when using Fakestack.

## Installation Issues

### Binary Not Found

**Error:**
```
FileNotFoundError: Binary not found
Error: Binary not found
```

**Cause:** The pre-built binary is missing or not in the expected location.

**Solutions:**

1. **Reinstall the package:**
   ```bash
   # Python
   pip uninstall fakestack
   pip install fakestack
   
   # Node.js
   npm uninstall fakestack
   npm install fakestack
   ```

2. **Check binary location:**
   ```bash
   # Python
   python -c "import fakestack; print(fakestack.__file__)"
   ls -la <path-from-above>/bin/
   
   # Node.js
   ls -la node_modules/fakestack/bin/
   ```

3. **Verify your platform is supported:**
   - darwin-amd64 (macOS Intel)
   - darwin-arm64 (macOS Apple Silicon)
   - linux-amd64 (Linux 64-bit)
   - windows-amd64 (Windows 64-bit)
   - windows-386 (Windows 32-bit)

### Permission Denied

**Error:**
```
PermissionError: [Errno 13] Permission denied
```

**Cause:** Binary doesn't have execute permissions.

**Solution:**
```bash
# Python
chmod +x <python-site-packages>/fakestack/bin/fakestack-*

# Node.js
chmod +x node_modules/fakestack/bin/fakestack-*

# Homebrew
brew reinstall 0xdps/fakestack/fakestack
```

### Import Error (Python)

**Error:**
```
ImportError: cannot import name 'Fakestack' from 'fakestack'
```

**Solution:**
```bash
# Check Python version (3.8+ required)
python --version

# Reinstall
pip install --force-reinstall fakestack

# Clear cache
pip cache purge
pip install fakestack
```

---

## Schema Issues

### Invalid JSON

**Error:**
```
json.JSONDecodeError: Expecting value: line 1 column 1
```

**Cause:** Malformed JSON schema.

**Solutions:**

1. **Validate JSON syntax:**
   ```bash
   # Online: https://jsonlint.com/
   
   # CLI
   python -m json.tool schema.json
   ```

2. **Common JSON mistakes:**
   ```json
   // ❌ Trailing commas
   {
     "database": {...},
     "tables": [...],  // <- Remove this comma
   }
   
   // ❌ Single quotes
   {'database': 'sqlite'}  // Use double quotes
   
   // ❌ Comments (not allowed in JSON)
   {
     "database": {...}  // This is a comment
   }
   ```

### Missing Required Fields

**Error:**
```
Error: database configuration required
Error: tables array required
```

**Cause:** Schema is missing required sections.

**Solution:**
```json
{
  "database": {  // ✅ Required
    "dbtype": "sqlite",
    "drivername": "sqlite",
    "database": "test.db"
  },
  "tables": [],  // ✅ Required (can be empty)
  "populate": []  // ✅ Required (can be empty)
}
```

### Invalid Column Type

**Error:**
```
Error: invalid column type
```

**Cause:** Unknown or incorrectly formatted column type.

**Solutions:**

```json
// ❌ Wrong
{"name": "age", "type": "number"}

// ✅ Correct
{"name": "age", "type": "integer"}

// ❌ String without length
{"name": "username", "type": "string"}

// ✅ String with length
{"name": "username", "type": {"name": "string", "args": {"length": 50}}}
```

### Foreign Key Reference Error

**Error:**
```
Error: foreign key table not found
Error: foreign key violation
```

**Cause:** Referenced table doesn't exist or isn't populated yet.

**Solutions:**

1. **Define tables in correct order:**
   ```json
   {
     "tables": [
       {"name": "users", ...},      // ✅ Parent first
       {"name": "posts", ...}       // ✅ Child second
     ]
   }
   ```

2. **Populate in correct order:**
   ```json
   {
     "populate": [
       {"name": "users", ...},      // ✅ Parent first
       {"name": "posts", ...}       // ✅ Child second
     ]
   }
   ```

---

## Database Issues

### SQLite

#### Database Locked

**Error:**
```
Error: database is locked
```

**Causes & Solutions:**

1. **Another process is using the database:**
   ```bash
   # Find processes using the file
   lsof database.db
   
   # Close other connections
   ```

2. **Increase busy timeout (schema modification):**
   ```json
   {
     "database": {
       "dbtype": "sqlite",
       "drivername": "sqlite",
       "database": "test.db"
     }
   }
   ```

3. **Delete and recreate:**
   ```bash
   rm database.db
   fakestack schema.json
   ```

#### File Permission Issues

**Error:**
```
Error: unable to open database file
```

**Solution:**
```bash
# Check directory permissions
ls -la

# Fix permissions
chmod 755 .
chmod 644 database.db
```

### MySQL

#### Connection Refused

**Error:**
```
Error: Can't connect to MySQL server
```

**Causes & Solutions:**

1. **MySQL not running:**
   ```bash
   # macOS
   brew services start mysql
   
   # Linux
   sudo systemctl start mysql
   
   # Check status
   brew services list  # macOS
   systemctl status mysql  # Linux
   ```

2. **Wrong port:**
   ```json
   {
     "database": {
       "port": 3306  // Default MySQL port
     }
   }
   ```

3. **Wrong host:**
   ```json
   {
     "database": {
       "host": "localhost"  // or "127.0.0.1"
     }
   }
   ```

#### Access Denied

**Error:**
```
Error: Access denied for user 'username'@'localhost'
```

**Solutions:**

1. **Check credentials:**
   ```bash
   mysql -u root -p
   # Enter password when prompted
   ```

2. **Create/grant permissions:**
   ```sql
   CREATE USER 'myuser'@'localhost' IDENTIFIED BY 'mypassword';
   GRANT ALL PRIVILEGES ON mydb.* TO 'myuser'@'localhost';
   FLUSH PRIVILEGES;
   ```

3. **Reset root password:**
   ```bash
   # macOS
   brew services stop mysql
   mysqld_safe --skip-grant-tables &
   mysql -u root
   # In MySQL prompt:
   ALTER USER 'root'@'localhost' IDENTIFIED BY 'newpassword';
   FLUSH PRIVILEGES;
   ```

#### Database Does Not Exist

**Error:**
```
Error: Unknown database 'mydb'
```

**Solution:**
```sql
CREATE DATABASE mydb;
```

### PostgreSQL

#### Connection Refused

**Error:**
```
Error: could not connect to server
```

**Solutions:**

1. **PostgreSQL not running:**
   ```bash
   # macOS
   brew services start postgresql
   
   # Linux
   sudo systemctl start postgresql
   
   # Check status
   brew services list  # macOS
   systemctl status postgresql  # Linux
   ```

2. **Wrong port:**
   ```json
   {
     "database": {
       "port": 5432  // Default PostgreSQL port
     }
   }
   ```

#### Authentication Failed

**Error:**
```
Error: password authentication failed
```

**Solutions:**

1. **Check pg_hba.conf:**
   ```bash
   # Find config file
   sudo -u postgres psql -c "SHOW hba_file;"
   
   # Edit file
   sudo nano /path/to/pg_hba.conf
   
   # Change to md5 for password auth:
   # local   all   all   md5
   # host    all   all   127.0.0.1/32   md5
   
   # Restart PostgreSQL
   sudo systemctl restart postgresql
   ```

2. **Reset password:**
   ```bash
   sudo -u postgres psql
   # In psql prompt:
   ALTER USER postgres WITH PASSWORD 'newpassword';
   ```

---

## Performance Issues

### Slow Generation

**Symptoms:** Takes too long to generate data.

**Solutions:**

1. **Reduce record count:**
   ```json
   {
     "populate": [
       {"name": "users", "count": 100}  // Start small
     ]
   }
   ```

2. **Use indexes:**
   ```json
   {
     "indexes": [
       {"name": "idx_user_id", "columns": ["user_id"], "unique": false}
     ]
   }
   ```

3. **Optimize database:**
   ```bash
   # SQLite
   sqlite3 database.db "VACUUM;"
   
   # MySQL
   mysqlcheck -o -u root -p mydb
   
   # PostgreSQL
   vacuumdb -U postgres mydb
   ```

4. **Batch processing:**
   Fakestack uses batches of 1000 records automatically.

### High Memory Usage

**Symptoms:** Process uses too much memory.

**Solutions:**

1. **Reduce batch size** (requires modifying Go core)
2. **Split into multiple schemas:**
   ```python
   # Generate users first
   users_schema = {...}
   Fakestack(users_schema).run()
   
   # Then posts
   posts_schema = {...}
   Fakestack(posts_schema).run()
   ```

---

## Platform-Specific Issues

### macOS

#### "Binary cannot be opened" Security Warning

**Error:**
```
"fakestack" cannot be opened because it is from an unidentified developer
```

**Solution:**
```bash
# Remove quarantine attribute
xattr -d com.apple.quarantine /path/to/fakestack/bin/fakestack-*

# Or via System Preferences:
# System Preferences > Security & Privacy > General > "Allow Anyway"
```

### Windows

#### PowerShell Execution Policy

**Error:**
```
cannot be loaded because running scripts is disabled
```

**Solution:**
```powershell
# Run PowerShell as Administrator
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

#### Path Length Limit

**Error:**
```
The filename or extension is too long
```

**Solution:**
```bash
# Enable long paths in Windows
# Run as Administrator:
reg add HKLM\SYSTEM\CurrentControlSet\Control\FileSystem /v LongPathsEnabled /t REG_DWORD /d 1 /f
```

### Linux

#### CGO / glibc Version

**Error:**
```
version 'GLIBC_X.XX' not found
```

**Solution:**
```bash
# Check glibc version
ldd --version

# If too old, upgrade system or use Docker
docker run -v $(pwd):/app -w /app python:3.11 \
  sh -c "pip install fakestack && fakestack schema.json"
```

---

## Getting Help

If you're still experiencing issues:

1. **Check existing issues:** [GitHub Issues](https://github.com/0xdps/fake-stack/issues)
2. **Create a new issue:** Include:
   - Operating system and version
   - Python/Node.js version
   - Complete error message
   - Minimal schema that reproduces the issue
3. **Community support:** [GitHub Discussions](https://github.com/0xdps/fake-stack/discussions)

## Common Debugging Steps

1. **Verify installation:**
   ```bash
   # Python
   pip show fakestack
   python -c "from fakestack import Fakestack; print('OK')"
   
   # Node.js
   npm list fakestack
   node -e "require('fakestack'); console.log('OK')"
   ```

2. **Test with minimal schema:**
   ```json
   {
     "database": {
       "dbtype": "sqlite",
       "drivername": "sqlite",
       "database": "test.db"
     },
     "tables": [],
     "populate": []
   }
   ```

3. **Enable verbose output:**
   ```python
   # Python
   import logging
   logging.basicConfig(level=logging.DEBUG)
   ```

4. **Check binary directly:**
   ```bash
   # Find binary
   find . -name "fakestack-*" -type f
   
   # Test execution
   ./path/to/binary --help
   ```

## See Also

- [Getting Started](getting-started.md) - Basic usage guide
- [Schema Reference](schema-reference.md) - Complete schema documentation
- [Database Support](databases.md) - Database-specific information
- [API Reference](api-reference.md) - Python and Node.js APIs
