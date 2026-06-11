"""Fakestack - High-Performance Database Generator

Generate databases from JSON schemas with realistic fake data.
Powered by a Go core for blazing-fast performance!

Auto-downloads latest binary on first run - no bundled binaries means tiny package size.
"""

__version__ = "2.0.0"
__author__ = "Devendra Pratap"
__email__ = "dps.manit@gmail.com"
__license__ = "MIT"

# Import main API
from .runner import fakestack, main, run_fakestack

__all__ = [
    "__version__",
    "__author__",
    "__email__",
    "main",
    "run_fakestack",
    "fakestack",
]
