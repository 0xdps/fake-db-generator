"""Integration tests for fakestack Go wrapper."""

import os
import tempfile
from pathlib import Path

import pytest


def test_import_fakestack():
    """Test that fakestack can be imported."""
    import fakestack
    
    assert hasattr(fakestack, '__version__')
    assert hasattr(fakestack, 'main')
    assert hasattr(fakestack, 'run_fakestack')
    assert fakestack.__version__ == "2.1.0"


def test_run_fakestack_help():
    """Test running fakestack with no arguments shows help."""
    from fakestack import run_fakestack
    
    # Should show error and usage
    exit_code = run_fakestack([])
    assert exit_code != 0


def test_download_schema():
    """Test downloading example schema."""
    from fakestack import run_fakestack
    
    with tempfile.TemporaryDirectory() as tmpdir:
        os.chdir(tmpdir)
        exit_code = run_fakestack(['-d', '.'])
        
        assert exit_code == 0
        assert Path('schema.json').exists()
        
        # Verify schema is valid JSON
        import json
        with open('schema.json') as f:
            schema = json.load(f)
        
        assert 'database' in schema
        assert 'tables' in schema
        assert 'populate' in schema


def test_create_and_populate():
    """Test creating tables and populating data."""
    from fakestack import run_fakestack
    
    with tempfile.TemporaryDirectory() as tmpdir:
        os.chdir(tmpdir)
        
        # Download schema
        exit_code = run_fakestack(['-d', '.'])
        assert exit_code == 0
        
        # Create tables and populate
        exit_code = run_fakestack(['-c', '-p', '-f', 'schema.json'])
        assert exit_code == 0
        
        # Verify database exists
        assert Path('test.db').exists()
        assert Path('test.db').stat().st_size > 0


def test_binary_detection():
    """Test that binary path detection works."""
    from fakestack.runner import get_binary_path
    
    binary_path = get_binary_path()
    
    assert binary_path.exists()
    assert binary_path.is_file()
    assert 'fakestack-' in binary_path.name


def test_version_consistency():
    """Test that version is consistent across package."""
    import fakestack
    
    # Check __init__.py version
    assert fakestack.__version__ == "2.1.0"
    
    # Check that version is a valid semantic version
    parts = fakestack.__version__.split('.')
    assert len(parts) == 3
    assert all(part.isdigit() for part in parts)


def test_cli_via_python_module():
    """Test running fakestack via python -m in same process."""
    from fakestack import run_fakestack
    
    with tempfile.TemporaryDirectory() as tmpdir:
        os.chdir(tmpdir)
        
        # Run via the imported function (simulates -m usage)
        exit_code = run_fakestack(['-d', '.'])
        
        assert exit_code == 0
        assert Path('schema.json').exists()


def test_multiple_runs():
    """Test that multiple runs work correctly."""
    from fakestack import run_fakestack
    
    with tempfile.TemporaryDirectory() as tmpdir:
        os.chdir(tmpdir)
        
        # First run
        exit_code = run_fakestack(['-d', '.'])
        assert exit_code == 0
        
        # Second run (should succeed or show already exists)
        exit_code = run_fakestack(['-d', '.'])
        # Either succeeds or file exists
        assert exit_code in [0, 1]


if __name__ == '__main__':
    pytest.main([__file__, '-v'])
