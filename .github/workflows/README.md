# Test Database Workflows

## Centralized Test Schema

The database integration tests use a centralized schema to ensure consistency across all supported databases.

### Schema File

**`test-schema.json`** - Contains the complete test schema with:
- **84 columns** covering all generator types
- **83 generator fields** (excluding auto-increment ID)
- All categories: Personal, Location, Network, Financial, Time, Localization, Product, Files, Media, Animals, Food, Vehicles, Identifiers, Text, and App data

### How It Works

Each database test job:
1. Merges the centralized schema with database-specific configuration using `jq`
2. Validates that exactly 83 generator fields are present
3. Creates and populates the test table
4. Verifies the row count

### Benefits

✅ **Single Source of Truth** - All generator fields defined once  
✅ **Consistency** - Same tests across all 6 databases  
✅ **Easy Updates** - Add new generators in one place  
✅ **Validation** - Automatic field count verification  
✅ **Reduced Duplication** - 71 lines vs 569 lines per database  

### Supported Databases

1. **MySQL** 8.0
2. **MariaDB** 10.11
3. **PostgreSQL** 14
4. **SQLite** 3.x
5. **MS SQL Server** 2022
6. **CockroachDB** Latest

### Adding New Generators

To add a new generator to all database tests:

1. Add the column to `test-schema.json` in the `tables[0].columns` array
2. Add the generator field to `test-schema.json` in the `populate[0].fields` array
3. Update the field count validation in `test-databases.yml` (increment the `83` in all 6 tests)

Example:
```json
// In tables[0].columns:
{"name": "new_field", "type": {"name": "string", "args": {"length": 50}}}

// In populate[0].fields:
{"name": "new_field", "generator": "new_generator"}
```

### Schema Validation

Each test automatically validates:
```bash
FIELD_COUNT=$(jq '.populate[0].fields | length' test-schema.json)
[ "$FIELD_COUNT" -eq 83 ] || (echo "❌ Expected 83 fields, got $FIELD_COUNT" && exit 1)
```

This ensures no fields are accidentally omitted during schema merging.
