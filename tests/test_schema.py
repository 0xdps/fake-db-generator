"""Tests for schema loading and validation."""

import json
import os
from pathlib import Path

import pytest

from fakestack.models.schema import (
    DbOptions,
    DbSchema,
    DbTable,
    DbType,
    TableColumn,
    load_schema,
)


class TestSchemaLoading:
    """Test schema loading functionality."""

    def test_load_example_schema(self):
        """Test loading example schema."""
        schema_path = (
            Path(__file__).parent.parent / "schemer" / "data" / "example-mysql.json"
        )
        if schema_path.exists():
            schema = load_schema(str(schema_path))
            assert isinstance(schema, DbSchema)
            assert schema.database.dbtype in [
                DbType.MySql,
                DbType.SqLite,
                DbType.Postgres,
            ]
            assert len(schema.tables) > 0

    def test_schema_validation(self):
        """Test schema model validation."""
        schema_dict = {
            "database": {
                "dbtype": "sqlite",
                "drivername": "sqlite",
                "username": "",
                "password": "",
                "host": "",
                "database": ":memory:",
            },
            "tables": [
                {
                    "name": "test_table",
                    "columns": [
                        {
                            "name": "id",
                            "type": "integer",
                            "options": {"primary_key": True},
                        }
                    ],
                    "indexes": [],
                }
            ],
            "populate": [],
        }

        schema = DbSchema(**schema_dict)
        assert schema.database.dbtype == DbType.SqLite
        assert len(schema.tables) == 1
        assert schema.tables[0].name == "test_table"


class TestTableColumn:
    """Test table column models."""

    def test_column_with_simple_type(self):
        """Test column with simple type string."""
        column = TableColumn(name="id", type="integer", options={"primary_key": True})
        assert column.name == "id"
        assert column.type.name == "integer"
        assert column.options["primary_key"] is True

    def test_column_with_complex_type(self):
        """Test column with complex type definition."""
        column = TableColumn(
            name="username",
            type={"name": "string", "args": {"length": 50}},
            options={"nullable": False},
        )
        assert column.name == "username"
        assert column.type.name == "string"
        assert column.type.args == {"length": 50}


class TestDbOptions:
    """Test database options model."""

    def test_sqlite_options(self):
        """Test SQLite database options."""
        options = DbOptions(
            dbtype=DbType.SqLite,
            drivername="sqlite",
            username="",
            password="",
            host="",
            database=":memory:",
        )
        assert options.dbtype == DbType.SqLite
        assert options.drivername == "sqlite"

    def test_mysql_options(self):
        """Test MySQL database options."""
        options = DbOptions(
            dbtype=DbType.MySql,
            drivername="mysql+mysqlconnector",
            username="root",
            password="password",
            host="localhost",
            database="testdb",
        )
        assert options.dbtype == DbType.MySql
        assert options.username == "root"
