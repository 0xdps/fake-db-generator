#!/usr/bin/env python3
"""Setup script for fakestack package."""

from setuptools import setup, find_packages
from pathlib import Path

# Read version from __init__.py
init_file = Path(__file__).parent / "fakestack" / "__init__.py"
version_line = [line for line in init_file.read_text().split('\n') if '__version__' in line][0]
version = version_line.split('=')[1].strip().strip('"')

setup(
    name="fakestack",
    version=version,
    packages=find_packages(),
    python_requires=">=3.8",
    install_requires=[],
    entry_points={
        "console_scripts": [
            "fakestack=fakestack.runner:main",
        ],
    },
)
