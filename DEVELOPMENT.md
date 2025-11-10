# Development Guide

Complete guide for developers working on Schemer.

## Prerequisites

- Python 3.8 or higher
- Git
- Virtual environment tool
- Database (MySQL, PostgreSQL, or SQLite for testing)

## Quick Start

```bash
# Clone and enter directory
git clone https://github.com/0xdps/fake-db-generator.git
cd fake-db-generator

# Set up virtual environment
python -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate

# Install in development mode
pip install -e ".[dev]"

# Verify setup
pytest
```

## Project Structure

```
fake-db-generator/
├── .github/
│   └── workflows/          # CI/CD workflows
│       ├── test-build.yml
│       ├── test-publish.yml
│       └── publish.yml
├── schemer/                # Main package
│   ├── __init__.py         # Version and public API
│   ├── runner.py           # CLI entry point
│   ├── data/               # Example schemas
│   └── models/             # Data models
│       ├── __init__.py
│       ├── fake.py         # Faker providers
│       ├── schema.py       # Pydantic models
│       └── utils.py        # Utilities
├── tests/                  # Test suite
│   ├── __init__.py
│   ├── test_schema.py
│   ├── test_faker.py
│   └── test_utils.py
├── pyproject.toml          # Package configuration
├── README.md
├── CONTRIBUTING.md         # Includes Code of Conduct & Security
├── DEVELOPMENT.md
├── CHANGELOG.md
└── LICENSE
```

## Development Workflow

### 1. Create Branch

```bash
git checkout -b feature/your-feature
```

### 2. Make Changes

Write code following project style guidelines (see below).

### 3. Write Tests

```bash
# Create test file
touch tests/test_your_feature.py

# Write tests (aim for >80% coverage)
```

Example test:

```python
import pytest
from schemer.models.schema import load_schema

def test_load_valid_schema():
    """Test loading a valid schema file."""
    schema = load_schema("schemer/data/example-mysql.json")
    assert schema.database.dbtype.value == "mysql"
    assert len(schema.tables) > 0
```

### 4. Run Tests

```bash
# All tests
pytest

# With coverage
pytest --cov=schemer --cov-report=html

# Specific test
pytest tests/test_schema.py::test_load_valid_schema

# Verbose output
pytest -v

# Show print statements
pytest -s
```

### 5. Code Quality

```bash
# Format code
black .
isort .

# Check style
flake8 .

# Type check
mypy schemer/
```

### 6. Commit

```bash
git add .
git commit -m "feat: add your feature"
```

Use [Conventional Commits](https://www.conventionalcommits.org/):
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `test`: Tests
- `refactor`: Code refactoring
- `style`: Formatting
- `chore`: Maintenance

### 7. Push and PR

```bash
git push origin feature/your-feature
# Create PR on GitHub
```

## Testing

### Running Tests

```bash
pytest                    # All tests
pytest -v                 # Verbose
pytest -x                 # Stop on first failure
pytest -k "schema"        # Run tests matching pattern
pytest --lf               # Run last failed
pytest -s                 # Show print statements
pytest --pdb              # Drop into debugger on failure
```

### Test Coverage

```bash
# Generate coverage report
pytest --cov=schemer --cov-report=html

# Open in browser
open htmlcov/index.html  # macOS
xdg-open htmlcov/index.html  # Linux
start htmlcov/index.html  # Windows
```

### Writing Tests

```python
import pytest
from schemer.models.fake import faker

class TestFaker:
    """Tests for custom Faker providers."""
    
    def test_person_generator(self):
        """Test person data generation."""
        person = faker.person()
        assert person.first_name
        assert person.last_name
        assert person.email
        assert person.gender in ["M", "F"]
    
    def test_random_from(self):
        """Test random selection from list."""
        options = ["a", "b", "c"]
        result = faker.random_from(*options)
        assert result in options
    
    @pytest.mark.parametrize("count", [1, 10, 100])
    def test_unique_generation(self, count):
        """Test unique value generation."""
        faker.unique.clear()
        values = [faker.unique.random_int() for _ in range(count)]
        assert len(values) == len(set(values))
```

### Test Fixtures

```python
@pytest.fixture
def sample_schema():
    """Provide sample schema for tests."""
    return {
        "database": {
            "dbtype": "sqlite",
            "drivername": "sqlite",
            "database": ":memory:"
        },
        "tables": [
            {
                "name": "test_table",
                "columns": [
                    {"name": "id", "type": "integer", "options": {"primary_key": True}}
                ]
            }
        ]
    }

def test_with_fixture(sample_schema):
    """Test using fixture."""
    assert sample_schema["database"]["dbtype"] == "sqlite"
```

## Code Style

### Python Guidelines

- **Formatting**: black (88 chars, default config)
- **Imports**: isort (black profile)
- **Linting**: flake8 (relaxed for line length)
- **Type Hints**: Add where beneficial, not required everywhere
- **Docstrings**: Google style for public APIs

### Docstring Example

```python
def generate_table_data(table_name: str, count: int) -> list:
    """Generate fake data for a database table.
    
    Creates the specified number of fake data rows for the given
    table using configured generators.
    
    Args:
        table_name: Name of the table to generate data for
        count: Number of rows to generate
        
    Returns:
        List of dictionaries containing generated data
        
    Raises:
        ValueError: If table_name is not found in schema
        
    Example:
        >>> data = generate_table_data("users", 10)
        >>> len(data)
        10
        >>> "username" in data[0]
        True
    """
    pass
```

### Import Organization

```python
# Standard library
import os
import sys
from pathlib import Path

# Third-party
from faker import Faker
from sqlalchemy import create_engine
from pydantic import BaseModel

# Local
from .models import DbSchema
from .utils import logger
```

## Debugging

### Using pdb

```python
import pdb; pdb.set_trace()  # Breakpoint
```

### Pytest Debugger

```bash
pytest --pdb              # Drop to pdb on failure
pytest -x --pdb           # Stop on first failure and debug
```

### Logging

```python
import logging

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

logger.debug("Debug information")
logger.info("General information")
logger.warning("Warning message")
logger.error("Error occurred")
```

## Building Package

### Local Build

```bash
# Clean previous builds
rm -rf build/ dist/ *.egg-info/

# Build package
python -m build

# Check package
twine check dist/*

# Test installation
pip install dist/*.whl
python -c "import schemer; print(schemer.__version__)"
```

### Test PyPI Upload (Optional)

```bash
# Upload to Test PyPI
twine upload --repository testpypi dist/*

# Install from Test PyPI
pip install --index-url https://test.pypi.org/simple/ schemer
```

## Release Process

### Version Bumping

1. Update version in `schemer/__init__.py`:
   ```python
   __version__ = "2.1.0"
   ```

2. Update `CHANGELOG.md`:
   ```markdown
   ## [2.1.0] - 2025-11-09
   
   ### Added
   - New feature X
   
   ### Fixed
   - Bug in Y
   ```

3. Commit and tag:
   ```bash
   git add schemer/__init__.py CHANGELOG.md
   git commit -m "chore: bump version to 2.1.0"
   git tag v2.1.0
   git push origin trunk
   git push origin v2.1.0
   ```

4. Create GitHub Release:
   - Go to GitHub → Releases → "Create a new release"
   - Select tag `v2.1.0`
   - Copy changelog content
   - Publish release (triggers PyPI publish)

## Common Tasks

### Add New Faker Provider

```python
# In schemer/models/fake.py
from faker.providers import BaseProvider

class MyProvider(BaseProvider):
    __provider__ = "my_provider"
    
    def my_generator(self, arg1: str) -> str:
        """Generate custom data."""
        return f"Generated: {arg1}"

# Register provider
faker.add_provider(MyProvider)
```

### Add New Column Type

```python
# In schemer/models/utils.py
from sqlalchemy import Boolean

def get_column_type(column: TableColumn):
    type_name = column.type.name
    type_args = column.type.args
    match type_name:
        case "boolean":
            return Boolean
        # ... existing cases
```

### Add Example Schema

```bash
# Create new schema file
cat > schemer/data/my-example.json << 'EOF'
{
  "database": {...},
  "tables": [...],
  "populate": [...]
}
EOF
```

## Troubleshooting

### Import Errors

```bash
# Reinstall package
pip install -e .
```

### Test Failures

```bash
# Clear cache
pytest --cache-clear
find . -type d -name __pycache__ -exec rm -rf {} +
```

### Environment Issues

```bash
# Reset environment
deactivate
rm -rf venv/
python -m venv venv
source venv/bin/activate
pip install -e ".[dev]"
```

### Database Connection Issues

```bash
# Test database connection
python -c "from sqlalchemy import create_engine; engine = create_engine('sqlite:///:memory:'); print('✅ Connection OK')"
```

## Resources

- [Python Packaging Guide](https://packaging.python.org/)
- [pytest Documentation](https://docs.pytest.org/)
- [black Documentation](https://black.readthedocs.io/)
- [SQLAlchemy Documentation](https://docs.sqlalchemy.org/)
- [Faker Documentation](https://faker.readthedocs.io/)
- [Pydantic Documentation](https://docs.pydantic.dev/)

## Getting Help

- **GitHub Discussions**: [Ask questions](https://github.com/0xdps/fake-db-generator/discussions)
- **GitHub Issues**: [Report bugs](https://github.com/0xdps/fake-db-generator/issues)
- **Email**: dps.manit@gmail.com

## Code Review Checklist

Before submitting PR:

- [ ] Code is formatted with `black`
- [ ] Imports sorted with `isort`
- [ ] No `flake8` errors
- [ ] Tests written and passing
- [ ] Coverage remains > 80%
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
- [ ] Commit messages follow convention
- [ ] PR description is clear

---

Happy coding! 🚀
