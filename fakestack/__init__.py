"""Schemer - JSON Schema to Database Generator

A Python tool for generating database tables and populating them with
realistic fake data based on JSON schema definitions.
"""

__version__ = "2.0.0"
__author__ = "Devendra Pratap"
__email__ = "dps.manit@gmail.com"
__license__ = "MIT"

from .models.fake import faker
from .models.schema import DbSchema, load_schema

# Import main components for public API
from .runner import base_setup, generate_tables, main, populate_data

__all__ = [
    "__version__",
    "__author__",
    "__email__",
    "main",
    "base_setup",
    "generate_tables",
    "populate_data",
    "DbSchema",
    "load_schema",
    "faker",
]
