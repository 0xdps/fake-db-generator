# Changelog

All notable changes to Fakestack will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Features in development

### Changed
- Changes to existing functionality

### Fixed
- Bug fixes

## [2.0.0] - 2025-11-09

### Added
- 🎉 **Major rebranding**: Renamed from `fake-db-generator` (fdg) to `Fakestack`
- 📦 Modern `pyproject.toml` configuration replacing `setup.py`
- 🔄 GitHub Actions CI/CD workflows for testing and publishing
- 🧪 Multi-platform testing (Ubuntu, Windows, macOS)
- 🐍 Python 3.8-3.13 compatibility testing
- 📝 Comprehensive documentation:
  - `CONTRIBUTING.md` - Contribution guidelines
  - `CODE_OF_CONDUCT.md` - Community standards
  - `SECURITY.md` - Security policy
  - `DEVELOPMENT.md` - Developer guide
  - `CHANGELOG.md` - Version history
- 🎨 Professional README with badges and better structure
- 🛠️ Development tools configuration (black, isort, flake8, mypy, pytest)
- ✅ Code quality checks in CI pipeline
- 📊 Test coverage reporting

### Changed
- ⚡ **Breaking**: Command changed from `fdg` to `fakestack`
- ⚡ **Breaking**: Package name changed from `fdg` to `fakestack`
- ⚡ **Breaking**: Import path changed from `fdg` to `fakestack`
- 📦 Updated dependencies to use modern version constraints
- 🔧 Improved package metadata and classifiers
- 📖 Enhanced documentation with emojis and better formatting
- 🎯 More descriptive help messages in CLI

### Fixed
- Package data files now properly included in distribution
- Import paths updated throughout codebase

### Migration Guide

If you were using version 1.x.x:

**Command Line:**
```bash
# Old
fdg -c -f schema.json

# New
fakestack -c -f schema.json
```

**Python Imports:**
```python
# Old
from fdg.models import DbSchema

# New  
from fakestack.models import DbSchema
```

**Installation:**
```bash
# Old
pip install fake-db-generator

# New
pip install fakestack
```

## [1.2.1] - 2024-XX-XX

### Fixed
- Minor bug fixes and improvements

## [1.2.0] - 2024-XX-XX

### Added
- Additional database support improvements
- Enhanced data generation capabilities

## [1.0.0] - 2024-XX-XX

### Added
- Initial stable release as `fake-db-generator`
- MySQL, PostgreSQL, and SQLite support
- JSON schema-based table generation
- Faker integration for realistic data
- Command-line interface
- Example schemas

[Unreleased]: https://github.com/0xdps/fake-db-generator/compare/v2.0.0...HEAD
[2.0.0]: https://github.com/0xdps/fake-db-generator/compare/v1.2.1...v2.0.0
[1.2.1]: https://github.com/0xdps/fake-db-generator/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/0xdps/fake-db-generator/compare/v1.0.0...v1.2.0
[1.0.0]: https://github.com/0xdps/fake-db-generator/releases/tag/v1.0.0
