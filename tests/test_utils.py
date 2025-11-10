"""Tests for utility functions."""

import pytest
from sqlalchemy import Date, ForeignKey, Integer, Numeric, String

from fakestack.models.schema import ColumnType, TableColumn
from fakestack.models.utils import get_column_type, ignore_exception, logger


class TestGetColumnType:
    """Test column type mapping."""

    def test_integer_type(self):
        """Test integer column type."""
        column = TableColumn(name="id", type="integer", options={})
        col_type = get_column_type(column)
        assert col_type is Integer

    def test_string_type_with_length(self):
        """Test string column type with length."""
        column = TableColumn(
            name="username", type={"name": "string", "args": {"length": 50}}, options={}
        )
        col_type = get_column_type(column)
        assert isinstance(col_type, type(String(50)))

    def test_date_type(self):
        """Test date column type."""
        column = TableColumn(name="created_at", type="date", options={})
        col_type = get_column_type(column)
        assert col_type is Date

    def test_numeric_type(self):
        """Test numeric column type."""
        column = TableColumn(
            name="price",
            type={"name": "number", "args": {"precision": 10, "scale": 2}},
            options={},
        )
        col_type = get_column_type(column)
        assert isinstance(col_type, type(Numeric(10, 2)))

    def test_foreign_key_type(self):
        """Test foreign key column type."""
        column = TableColumn(
            name="user_id", type={"name": "foreign", "args": "users.id"}, options={}
        )
        col_type = get_column_type(column)
        assert isinstance(col_type, ForeignKey)

    def test_invalid_type(self):
        """Test that invalid type raises exception."""
        column = TableColumn(name="invalid", type="invalid_type", options={})
        with pytest.raises(Exception, match="No matched column type found"):
            get_column_type(column)


class TestLogger:
    """Test logger function."""

    def test_logger_creation(self):
        """Test creating a logger."""
        log_func = logger(10, 5)
        assert callable(log_func)

    def test_logger_formats_output(self, capsys):
        """Test logger output formatting."""
        log_func = logger(10, 3)

        # Log some progress
        log_func("test_table", 1, 100)
        captured = capsys.readouterr()
        assert "test_table" in captured.out
        assert "1" in captured.out


class TestIgnoreException:
    """Test ignore_exception context manager."""

    def test_ignore_specified_exception(self):
        """Test that specified exceptions are ignored."""
        with ignore_exception(ValueError):
            raise ValueError("This should be ignored")
        # If we get here, the exception was ignored successfully

    def test_do_not_ignore_other_exceptions(self):
        """Test that other exceptions are not ignored."""
        with pytest.raises(TypeError):
            with ignore_exception(ValueError):
                raise TypeError("This should not be ignored")

    def test_ignore_multiple_exceptions(self):
        """Test ignoring multiple exception types."""
        with ignore_exception(ValueError, TypeError):
            raise ValueError("Ignored")

        with ignore_exception(ValueError, TypeError):
            raise TypeError("Also ignored")

    def test_no_exception_raised(self):
        """Test normal operation when no exception is raised."""
        result = None
        with ignore_exception(Exception):
            result = "success"
        assert result == "success"
