#!/bin/bash

# Fakestack Uninstall Script
# Removes fakestack from npm, pip, and Homebrew

set -e

echo "🗑️  Fakestack Uninstall Script"
echo "=============================="
echo ""

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Track if anything was uninstalled
UNINSTALLED=false

# Uninstall from npm (global)
echo "📦 Checking npm (global)..."
if npm list -g fakestack 2>/dev/null | grep -q fakestack; then
    echo -e "${YELLOW}Uninstalling fakestack from npm...${NC}"
    npm uninstall -g fakestack
    echo -e "${GREEN}✓ Uninstalled from npm${NC}"
    UNINSTALLED=true
else
    echo "  Not installed globally in npm"
fi
echo ""

# Uninstall from pip (current environment)
echo "🐍 Checking pip (current environment)..."
if pip show fakestack &>/dev/null; then
    echo -e "${YELLOW}Uninstalling fakestack from pip...${NC}"
    pip uninstall -y fakestack
    echo -e "${GREEN}✓ Uninstalled from pip${NC}"
    UNINSTALLED=true
else
    echo "  Not installed in current pip environment"
fi
echo ""

# Uninstall from Homebrew
echo "🍺 Checking Homebrew..."
if brew list fakestack &>/dev/null; then
    echo -e "${YELLOW}Uninstalling fakestack from Homebrew...${NC}"
    brew uninstall fakestack
    echo -e "${GREEN}✓ Uninstalled from Homebrew${NC}"
    UNINSTALLED=true
else
    echo "  Not installed in Homebrew"
fi
echo ""

# Optional: Remove tap
if brew tap | grep -q "0xdps/packages"; then
    echo "🔧 Homebrew tap 0xdps/packages is still installed"
    read -p "Do you want to remove the tap? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        brew untap 0xdps/packages
        echo -e "${GREEN}✓ Removed tap 0xdps/packages${NC}"
    fi
    echo ""
fi

# Check pyenv for other Python versions
echo "🔍 Checking other Python environments..."
if command -v pyenv &>/dev/null; then
    PYTHON_VERSIONS=$(pyenv versions --bare 2>/dev/null || echo "")
    if [ -n "$PYTHON_VERSIONS" ]; then
        echo "  Found pyenv Python versions. Checking each..."
        while IFS= read -r version; do
            if pyenv shell "$version" 2>/dev/null && pip show fakestack &>/dev/null; then
                echo -e "  ${YELLOW}Found in Python $version${NC}"
                read -p "  Uninstall from Python $version? (y/N): " -n 1 -r
                echo
                if [[ $REPLY =~ ^[Yy]$ ]]; then
                    pip uninstall -y fakestack
                    echo -e "  ${GREEN}✓ Uninstalled from Python $version${NC}"
                    UNINSTALLED=true
                fi
            fi
        done <<< "$PYTHON_VERSIONS"
        pyenv shell --unset 2>/dev/null || true
    fi
fi
echo ""

# Summary
echo "=============================="
if [ "$UNINSTALLED" = true ]; then
    echo -e "${GREEN}✓ Uninstall complete!${NC}"
else
    echo -e "${YELLOW}No installations found${NC}"
fi
echo ""
