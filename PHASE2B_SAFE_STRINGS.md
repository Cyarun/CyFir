# Phase 2B: Safe String Updates Plan

## Objective
Update user-visible strings that DO NOT affect:
- Protocol compatibility
- Data formats
- API contracts
- Service names
- File formats

## Categories of SAFE String Updates

### 1. Log Messages (SAFE if not parsed)
**Safe to update:**
- Info/debug log messages
- Error messages shown to users
- Status messages

**NOT safe:**
- Log messages that are parsed by monitoring tools
- Error codes that are checked programmatically

### 2. Help Text and Descriptions (SAFE)
- Command help text
- Flag descriptions
- User-facing documentation strings

### 3. Display Names (SAFE with care)
- Application title in GUI
- Menu items
- Button labels

### 4. Comments in Code (ALWAYS SAFE)
- File headers
- Function comments
- TODO comments

## String Update Priority Order

### Phase 2B-1: Super Safe Updates
1. Copyright headers in files
2. Code comments
3. Help text for commands
4. Log messages that are clearly informational

### Phase 2B-2: Medium Safe Updates  
1. Error messages shown to users
2. GUI labels and text
3. Status messages

### Phase 2B-3: Careful Updates
1. Service display names (NOT service IDs)
2. Configuration field descriptions
3. Artifact descriptions

## Strings to NEVER Update
❌ Service registry names
❌ Protocol magic strings
❌ API endpoint paths
❌ Database table/column names
❌ File format identifiers
❌ Certificate CN/SAN fields
❌ Artifact namespaces

## Implementation Strategy

### Step 1: Find Safe Candidates
```bash
# Find log messages
grep -r "Velociraptor" --include="*.go" | grep -E "(log\.|Log\(|fmt\.Printf)" 

# Find help text
grep -r "Velociraptor" --include="*.go" | grep -i "help\|description"

# Find comments
grep -r "Velociraptor" --include="*.go" | grep "//"
```

### Step 2: Test After Each Change
1. Make small batch of changes
2. Compile and test
3. Run smoke tests
4. Commit if successful

### Step 3: Validation
- Binary still starts
- Help commands work
- No parsing errors
- GUI loads correctly

## First Safe Targets

Let's start with the safest possible changes...