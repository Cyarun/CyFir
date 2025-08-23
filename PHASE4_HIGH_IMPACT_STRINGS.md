# Phase 4: High-Impact User-Visible Strings

## Strategy
Focus on strings that users actually see during normal operation, skipping internal development strings.

## Target Categories (Priority Order)

### 1. Error Messages (User-Facing)
- Login errors
- Connection errors
- Configuration errors
- Permission errors

### 2. Status Messages
- Server startup messages
- Client connection messages
- Operation status updates

### 3. GUI Text
- Window titles
- Menu items
- Button labels
- Status bar text

### 4. Command Output
- Command success messages
- Progress indicators
- Result summaries

## Search Patterns
- Error messages: `fmt.Errorf`, `errors.New`, `return.*Error`
- User messages: `fmt.Printf`, `fmt.Println`, `scope.Log`
- GUI strings: Labels, titles, descriptions
- Status updates: "Starting", "Connected", "Running"