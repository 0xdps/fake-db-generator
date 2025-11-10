# Contributing to Fakestack

Thank you for your interest in contributing to Fakestack! This guide will help you get started.

## 📋 Ways to Contribute

- 🐛 Report bugs
- 💡 Suggest features
- 📖 Improve documentation
- 🔧 Submit bug fixes
- ✨ Add new features
- 🧪 Write tests
- 🎨 Improve code quality

## 🚀 Getting Started

### 1. Fork and Clone

```bash
git clone https://github.com/YOUR-USERNAME/fake-stack.git
cd fake-stack
```

### 2. Set Up Development Environment

Choose the language you're working with:

**Go Core:**
```bash
cd golang
go mod download
go build
go test -v ./...
```

**Python Wrapper:**
```bash
cd python
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
pip install -e ".[dev]"
pytest tests/ -v
```

**Node.js Wrapper:**
```bash
cd node
npm install
npm run build
npm test
```

### 3. Create a Branch

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b bugfix/fix-issue-123
```

## 🧪 Testing

### Go Tests
```bash
cd golang
go test -v ./...
go test -race ./...
go test -cover ./...
```

### Python Tests
```bash
cd python
pytest tests/ -v
pytest --cov=fakestack --cov-report=html
black --check fakestack/
```

### Node.js Tests
```bash
cd node
npm test
npm run lint
npm run build
```

### Integration Tests

All wrappers should be tested to ensure they properly execute the Go binary:
```bash
# Python integration test
cd python && pytest tests/ -v

# Node.js integration test
cd node && npm test
```

## ✏️ Making Changes

### Commit Message Format

Use conventional commits:

```
type(scope): brief description

Detailed explanation if needed.

Fixes #123
```

**Types:** `feat`, `fix`, `docs`, `test`, `refactor`, `style`, `chore`

**Examples:**
```
feat(generators): add new date generator
fix(schema): handle empty column lists
docs(readme): update installation instructions
test(faker): add tests for person provider
```

## 📤 Submitting Changes

### Pull Request Process

1. **Update your branch**
   ```bash
   git fetch upstream
   git rebase upstream/trunk
   ```

2. **Run all checks**
   ```bash
   pytest
   black .
   isort .
   flake8 .
   ```

3. **Push and create PR**
   ```bash
   git push origin feature/your-feature
   ```

4. **Fill out PR template** with description and context

### PR Guidelines

- ✅ One feature/fix per PR
- ✅ Include tests
- ✅ Update documentation
- ✅ Pass all CI checks
- ✅ Address review comments

## 🎨 Code Style

- Use `black` for formatting (88 char line length)
- Use `isort` for import sorting
- Follow PEP 8 guidelines
- Add type hints where beneficial
- Write docstrings for public APIs

### Docstring Example

```python
def generate_data(count: int, generator: str) -> list:
    """Generate fake data using specified generator.
    
    Args:
        count: Number of items to generate
        generator: Name of the Faker generator to use
        
    Returns:
        List of generated data items
        
    Raises:
        ValueError: If generator is not found
        
    Example:
        >>> generate_data(5, "name")
        ['John Doe', 'Jane Smith', ...]
    """
    pass
```

## 📚 Documentation

Update relevant documentation:
- Docstrings in code
- README.md for user-facing changes
- Tutorials for new features
- CHANGELOG.md

## 🐛 Reporting Bugs

Include:
- Clear description
- Steps to reproduce
- Expected vs actual behavior
- Environment (Python version, OS, database)
- Schema file (if applicable)
- Minimal code example

**Bug Report Template:**

```markdown
**Description:** Brief description of the bug

**Steps to Reproduce:**
1. Run `fakestack -c -f schema.json`
2. See error

**Expected Behavior:** What should happen

**Actual Behavior:** What actually happens

**Environment:**
- OS: macOS 14.0
- Python: 3.11.0
- Fakestack: 2.0.0
- Database: MySQL 8.0

**Schema File:**
```json
{...}
```

**Error Message:**
```
Error traceback here
```
```

## 💡 Suggesting Features

Include:
- Use case and problem you're solving
- Proposed solution
- Alternative approaches considered
- Code examples showing usage

**Feature Request Template:**

```markdown
**Problem:** Description of the problem

**Proposed Solution:** How to solve it

**Example Usage:**
```python
# Show how the feature would be used
```

**Alternatives Considered:**
- Alternative 1: ...
- Alternative 2: ...

**Additional Context:** Any other relevant information
```

## 🔍 Development Tips

### Running Specific Tests

**Go:**
```bash
cd golang
go test -v -run TestSchemaParser
go test -v ./generator/...
```

**Python:**
```bash
cd python
pytest tests/test_schema.py -v
pytest tests/test_schema.py::test_load_schema
pytest -k "schema"
```

**Node.js:**
```bash
cd node
npm test -- --testNamePattern="schema"
```

### Debugging

**Go:**
```bash
cd golang
go test -v -run TestYourTest
dlv test -- -test.run TestYourTest
```

**Python:**
```bash
cd python
pytest -v -s
pytest --pdb
```

**Node.js:**
```bash
cd node
npm test -- --verbose
node --inspect-brk node_modules/.bin/jest
```

### Local Package Build

**Python:**
```bash
cd python
python -m build
twine check dist/*
pip install dist/*.whl
```

**Node.js:**
```bash
cd node
npm run build
npm pack
npm install -g fakestack-*.tgz
```

**Go Binary:**
```bash
cd golang
go build -o fakestack
./fakestack --help
```

## 📝 Documentation Standards

- Keep README.md up to date
- Add docstrings to all public functions/classes
- Update CHANGELOG.md with your changes
- Create tutorials for major features
- Add examples to the examples/ directory

## 📁 Project Structure

```
fake-stack/
├── golang/          # Go core implementation
│   ├── *.go        # Source files (main.go, schema.go, etc.)
│   ├── go.mod      # Go dependencies
│   └── README.md   # Go-specific docs
├── python/          # Python wrapper (PyPI: fakestack)
│   ├── fakestack/  # Python module
│   │   ├── __init__.py
│   │   ├── runner.py
│   │   └── bin/    # Pre-built binaries
│   ├── tests/      # Python integration tests
│   ├── pyproject.toml
│   └── README.md   # Python-specific docs
├── node/            # Node.js wrapper (npm: fakestack)
│   ├── src/        # TypeScript source
│   ├── dist/       # Compiled JavaScript
│   ├── bin/        # Pre-built binaries
│   ├── tests/      # Node.js integration tests
│   ├── package.json
│   └── README.md   # Node-specific docs
├── bin/             # Compiled binaries for all platforms
├── scripts/         # Build and release scripts
└── .github/         # CI/CD workflows
```

## ❓ Questions?

- Open a [Discussion](https://github.com/0xdps/fake-stack/discussions)
- Check [Documentation](https://github.com/0xdps/fake-stack#readme)
- Email: dps.manit@gmail.com

## 🙏 Thank You!

Your contributions make this project better! Every contribution, no matter how small, is valued and appreciated.

## 📜 Code of Conduct

We are committed to providing a welcoming and inclusive environment for everyone. By participating in this project, you agree to:

### Our Standards

**Positive behavior:**
- Use welcoming and inclusive language
- Be respectful of differing viewpoints
- Accept constructive criticism gracefully
- Focus on what's best for the community
- Show empathy towards others

**Unacceptable behavior:**
- Harassment, trolling, or insulting comments
- Personal or political attacks
- Publishing others' private information
- Any conduct inappropriate in a professional setting

### Enforcement

Instances of unacceptable behavior may be reported to **dps.manit@gmail.com**. All complaints will be reviewed and investigated promptly and fairly.

## 🔒 Security Policy

### Reporting Vulnerabilities

**Do not report security vulnerabilities through public GitHub issues.**

To report security issues, email **dps.manit@gmail.com** with:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

You should receive a response within 48 hours.

### Security Best Practices

When contributing:
- Never commit secrets, API keys, or passwords
- Validate all input from JSON schemas
- Use parameterized queries for database operations
- Keep dependencies up to date

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.
