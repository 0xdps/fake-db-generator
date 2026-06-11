"""
Fakestack - High-Performance Database Generator

Python wrapper for the Go core binary.
Downloads binary on first run and auto-updates when new versions are available.
"""

import json
import os
import platform
import subprocess
import sys
import urllib.request
from pathlib import Path


def get_cache_dir():
    """Get the fakestack cache directory."""
    if platform.system() == "Windows":
        cache_base = os.getenv("LOCALAPPDATA", os.path.expanduser("~"))
    else:
        cache_base = os.getenv("XDG_CACHE_HOME", os.path.expanduser("~/.cache"))
    
    cache_dir = Path(cache_base) / "fakestack" / "bin"
    cache_dir.mkdir(parents=True, exist_ok=True)
    return cache_dir


def get_version_file():
    """Get path to version tracking file."""
    if platform.system() == "Windows":
        cache_base = os.getenv("LOCALAPPDATA", os.path.expanduser("~"))
    else:
        cache_base = os.getenv("XDG_CACHE_HOME", os.path.expanduser("~/.cache"))
    
    version_file = Path(cache_base) / "fakestack" / "version.txt"
    version_file.parent.mkdir(parents=True, exist_ok=True)
    return version_file


def get_platform_info():
    """
    Detect platform and architecture.

    Returns:
        tuple: (os_name, arch_name, binary_name)
    """
    # Detect operating system
    system = platform.system().lower()
    system_map = {"linux": "linux", "darwin": "darwin", "windows": "windows"}
    os_name = system_map.get(system, system)

    # Detect architecture
    machine = platform.machine().lower()
    arch_map = {
        "x86_64": "amd64",
        "amd64": "amd64",
        "arm64": "arm64",
        "aarch64": "arm64",
    }
    arch_name = arch_map.get(machine, "amd64")

    # Construct binary name
    binary_name = f"fakestack-{os_name}-{arch_name}"
    if os_name == "windows":
        binary_name += ".exe"

    return os_name, arch_name, binary_name


def get_latest_version():
    """
    Fetch latest version from GitHub releases.

    Returns:
        str: Latest version number (e.g., "1.2.0")
    """
    try:
        url = "https://api.github.com/repos/0xdps/fake-stack/releases/latest"
        headers = {"User-Agent": "fakestack-python"}
        
        req = urllib.request.Request(url, headers=headers)
        with urllib.request.urlopen(req, timeout=10) as response:
            data = json.loads(response.read().decode())
            version = data.get("tag_name", "v1.2.0").lstrip("v")
            return version
    except Exception as e:
        # If we can't fetch, return a default version
        print(f"Warning: Could not check for updates: {e}", file=sys.stderr)
        return None


def download_binary(version, binary_name, target_path):
    """
    Download binary from GitHub releases.

    Args:
        version (str): Version to download
        binary_name (str): Name of the binary file
        target_path (Path): Where to save the binary
    """
    print(f"📦 Downloading fakestack v{version}...", file=sys.stderr)
    
    url = f"https://github.com/0xdps/fake-stack/releases/download/v{version}/{binary_name}"
    
    try:
        # Download with progress
        urllib.request.urlretrieve(url, target_path)
        
        # Make executable on Unix systems
        if platform.system() != "Windows":
            os.chmod(target_path, 0o755)
        
        print("✓ Download complete!", file=sys.stderr)
    except Exception as e:
        if target_path.exists():
            target_path.unlink()
        raise RuntimeError(f"Failed to download binary: {e}")


def ensure_binary():
    """
    Ensure binary is downloaded and up-to-date.

    Returns:
        Path: Path to the binary
    """
    _, _, binary_name = get_platform_info()
    cache_dir = get_cache_dir()
    binary_path = cache_dir / binary_name
    version_file = get_version_file()
    
    # Check if binary exists
    binary_exists = binary_path.exists()
    
    # Get latest version
    latest_version = get_latest_version()
    
    # If we can't check version but binary exists, use it
    if latest_version is None:
        if binary_exists:
            return binary_path
        raise RuntimeError(
            "Failed to check for latest version and no local binary found. "
            "Please check your internet connection."
        )
    
    # Check local version
    local_version = ""
    if version_file.exists():
        try:
            local_version = version_file.read_text().strip()
        except Exception:
            pass
    
    # Download if missing or outdated
    if not binary_exists or local_version != latest_version:
        if binary_exists and local_version != latest_version:
            print(f"🔄 Updating from v{local_version} to v{latest_version}...", file=sys.stderr)
        
        download_binary(latest_version, binary_name, binary_path)
        
        # Save version
        version_file.write_text(latest_version)
    
    return binary_path


def get_binary_path():
    """
    Get path to the fakestack binary (legacy, for compatibility).

    Returns:
        Path: Path to the platform-specific fakestack binary

    Raises:
        RuntimeError: If binary cannot be obtained
    """
    return ensure_binary()


def run_fakestack(args=None):
    """
    Execute the Go binary with given arguments.

    Args:
        args (list, optional): Command-line arguments. Defaults to sys.argv[1:].

    Returns:
        int: Exit code from the Go binary (0 = success, non-zero = error)
    """
    if args is None:
        args = sys.argv[1:]

    try:
        # Ensure binary is downloaded and up-to-date
        binary = ensure_binary()

        # Execute the Go binary
        result = subprocess.run(
            [str(binary)] + args,
            cwd=os.getcwd(),
            # Stream output directly to console
            stdout=None,
            stderr=None,
        )

        return result.returncode

    except RuntimeError as e:
        print(f"Error: {e}", file=sys.stderr)
        return 1
    except KeyboardInterrupt:
        print("\nInterrupted by user", file=sys.stderr)
        return 130
    except Exception as e:
        print(f"Unexpected error: {e}", file=sys.stderr)
        return 1


def fakestack(args=None):
    """
    Main API entry point. Alias for run_fakestack.

    Args:
        args (list, optional): Command-line arguments

    Returns:
        int: Exit code (0 = success)
    """
    return run_fakestack(args)


def main():
    """CLI entry point for fakestack command."""
    exit_code = run_fakestack()
    sys.exit(exit_code)


if __name__ == "__main__":
    main()
