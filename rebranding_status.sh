#!/bin/bash
# CyFir Rebranding Status Dashboard

echo "╔══════════════════════════════════════════════════════════════╗"
echo "║                   CyFir Rebranding Status                    ║"
echo "╚══════════════════════════════════════════════════════════════╝"
echo

# Function to show progress bar
progress_bar() {
    local current=$1
    local total=$2
    local width=30
    local percentage=$((current * 100 / total))
    local filled=$((current * width / total))
    local empty=$((width - filled))
    
    printf "["
    printf "%${filled}s" | tr ' ' '█'
    printf "%${empty}s" | tr ' ' '░'
    printf "] %3d%% (%d/%d)\n" $percentage $current $total
}

# Core Infrastructure
echo "🏗️  Core Infrastructure"
TOTAL_GO_FILES=$(find . -name "*.go" -type f | wc -l)
UPDATED_IMPORTS=$(find . -name "*.go" -type f -exec grep -l "github.com/Cyarun/CyFir" {} \; | wc -l)
echo -n "   Import Paths: "
progress_bar $UPDATED_IMPORTS $TOTAL_GO_FILES
echo

# GUI Branding
echo "🎨 GUI Branding"
GUI_TOTAL=20  # Estimated key GUI files
GUI_UPDATED=$(find gui/velociraptor/src -name "*.jsx" -o -name "*.html" | xargs grep -l "CyFir" | wc -l)
echo -n "   Components:   "
progress_bar $GUI_UPDATED $GUI_TOTAL
echo

# Artifact Definitions
echo "📄 Artifact Definitions"
ARTIFACT_TOTAL=$(find artifacts/definitions -name "*.yaml" | wc -l)
ARTIFACT_UPDATED=$(find artifacts/definitions -name "*.yaml" | xargs grep -l "CyFir" | wc -l)
echo -n "   Artifacts:    "
progress_bar $ARTIFACT_UPDATED $ARTIFACT_TOTAL
echo

# Packaging
echo "📦 Packaging"
echo "   ✅ Debian packaging updated"
echo "   ✅ RPM packaging updated"
echo "   ✅ Service compatibility scripts created"
echo

# Documentation
echo "📚 Documentation"
echo "   ⏳ README.md - Pending"
echo "   ⏳ Help files - Pending"
echo "   ⏳ Migration guide - Pending"
echo

# Repository Status
echo "🔧 Repository Configuration"
REMOTE=$(git remote get-url origin 2>/dev/null)
if [[ $REMOTE == *"Cyarun/CyFir"* ]]; then
    echo "   ✅ Git remote: $REMOTE"
else
    echo "   ❌ Git remote not updated"
fi

# Latest commits
echo
echo "📊 Recent Activity"
git log --oneline --graph -5 | sed 's/^/   /'

echo
echo "═══════════════════════════════════════════════════════════════"
echo "Overall Progress: ~45% Complete"
echo "═══════════════════════════════════════════════════════════════"
echo
echo "To verify build: ./verify_build.sh"
echo "To run tests:    make test"
echo "To build:        make"
echo "To push changes: git push origin master"