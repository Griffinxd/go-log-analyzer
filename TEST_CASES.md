# Test Cases for Log Analyzer

## Test Files Overview

### 1. `app_normal.log` - Basic Functionality
**Purpose:** General mixed entries with all levels represented

**Contains:**
- 10 total entries
- 4 INFO entries
- 2 WARNING entries
- 4 ERROR entries
- Timestamp range: 10:30:45 to 10:37:00

**Test Cases:**

#### Test 1.1: No filters - print all
```bash
go run main.go -file=app_normal.log
```
Expected: All 10 entries printed in text format

#### Test 1.2: Filter by ERROR level
```bash
go run main.go -file=app_normal.log -level=ERROR
```
Expected: 4 entries (all ERROR level)
```
[2024-01-15 10:33:00] ERROR Database connection failed
[2024-01-15 10:33:05] ERROR Retrying connection...
[2024-01-15 10:36:00] ERROR Invalid request format
[2024-01-15 10:37:00] ERROR Timeout on external API call
```

#### Test 1.3: Filter by INFO level
```bash
go run main.go -file=app_normal.log -level=INFO
```
Expected: 4 entries (all INFO level)

#### Test 1.4: Statistics
```bash
go run main.go -file=app_normal.log -stats
```
Expected output format:
```
Total: 10 | INFO: 4 | WARNING: 2 | ERROR: 4 | Most Common Error: Database connection failed
```

#### Test 1.5: JSON format
```bash
go run main.go -file=app_normal.log -level=WARNING -format=json
```
Expected: 2 JSON objects (one per line)
```
{"timestamp":"2024-01-15T10:32:15Z","level":"WARNING","message":"High memory usage detected"}
{"timestamp":"2024-01-15T10:35:22Z","level":"WARNING","message":"Cache miss rate 45%"}
```

#### Test 1.6: Time range filter (inclusive)
```bash
go run main.go -file=app_normal.log -from="2024-01-15 10:33:00" -to="2024-01-15 10:36:00"
```
Expected: 5 entries within that range (inclusive)

---

### 2. `app_edge_cases.log` - Error Handling
**Purpose:** Test robustness with malformed entries

**Contains:**
- Mix of valid and invalid entries
- Missing brackets, invalid timestamps, missing levels
- Case variations (lowercase "info")
- Entries with special characters in message

**Test Cases:**

#### Test 2.1: No crashes on malformed data
```bash
go run main.go -file=app_edge_cases.log
```
Expected: Program runs without crashing, valid entries printed, invalid entries skipped

#### Test 2.2: Lowercase levels work
```bash
go run main.go -file=app_edge_cases.log -level=INFO
```
Expected: Should find both "INFO" and "info" entries (case-insensitive)

#### Test 2.3: Messages with brackets
```bash
go run main.go -file=app_edge_cases.log
```
Expected: Entry "Message with [brackets] inside" parses correctly

---

### 3. `app_errors_only.log` - Statistics Testing
**Purpose:** Test most common error detection

**Contains:**
- 10 ERROR entries only
- "Database connection failed" appears 5 times (most common)
- Other errors appear once or twice

**Test Cases:**

#### Test 3.1: Most common error detection
```bash
go run main.go -file=app_errors_only.log -stats
```
Expected output:
```
Total: 10 | INFO: 0 | WARNING: 0 | ERROR: 10 | Most Common Error: Database connection failed
```

#### Test 3.2: Filter errors and count
```bash
go run main.go -file=app_errors_only.log -level=ERROR
```
Expected: All 10 entries printed

---

### 4. `app_large.log` - Volume Testing
**Purpose:** Test with more realistic log volume

**Contains:**
- 30 total entries
- 16 INFO entries
- 7 WARNING entries
- 7 ERROR entries
- Time span: 08:00:00 to 08:15:00

**Test Cases:**

#### Test 4.1: Full output
```bash
go run main.go -file=app_large.log
```
Expected: All 30 entries printed

#### Test 4.2: Filter by level
```bash
go run main.go -file=app_large.log -level=WARNING
```
Expected: 7 WARNING entries

#### Test 4.3: Statistics on large file
```bash
go run main.go -file=app_large.log -stats
```
Expected:
```
Total: 30 | INFO: 16 | WARNING: 7 | ERROR: 7 | Most Common Error: Job failed: timeout
```
(or another error if frequency is different)

#### Test 4.4: Time range from first entry
```bash
go run main.go -file=app_large.log -from="2024-01-15 08:00:00" -to="2024-01-15 08:05:00"
```
Expected: 6 entries from 08:00:00 to 08:05:00

#### Test 4.5: JSON output on large file
```bash
go run main.go -file=app_large.log -level=ERROR -format=json
```
Expected: 7 JSON objects (all ERROR entries)

---

### 5. `app_time_range.log` - Time Filter Testing
**Purpose:** Specifically test timestamp filtering

**Contains:**
- Entries from 09:00 to 13:00
- 3 entries before range, 7 in range, 2 after range
- Range for testing: 10:00:00 to 12:00:00

**Test Cases:**

#### Test 5.1: Time range (inclusive both ends)
```bash
go run main.go -file=app_time_range.log -from="2024-01-15 10:00:00" -to="2024-01-15 12:00:00"
```
Expected: 9 entries (including both boundary timestamps)
- Start at "Start of filter range" (10:00:00)
- End at "End of filter range" (12:00:00)

#### Test 5.2: From only (no upper bound)
```bash
go run main.go -file=app_time_range.log -from="2024-01-15 10:30:00"
```
Expected: 8 entries from 10:30:00 onwards

#### Test 5.3: To only (no lower bound)
```bash
go run main.go -file=app_time_range.log -to="2024-01-15 11:00:00"
```
Expected: 7 entries up to 11:00:00

#### Test 5.4: Time range with level filter
```bash
go run main.go -file=app_time_range.log -from="2024-01-15 10:00:00" -to="2024-01-15 12:00:00" -level=ERROR
```
Expected: 2 ERROR entries within range

---

### 6. `app_empty.log` - Empty File
**Purpose:** Test error handling for empty files

**Test Cases:**

#### Test 6.1: Empty file handling
```bash
go run main.go -file=app_empty.log
```
Expected: No output (or graceful message) - no crash

#### Test 6.2: Empty file with stats
```bash
go run main.go -file=app_empty.log -stats
```
Expected:
```
Total: 0 | INFO: 0 | WARNING: 0 | ERROR: 0 | Most Common Error: 
```

---

## Additional Test Scenarios

### Test 7: File not found
```bash
go run main.go -file=nonexistent.log
```
Expected error: "Error: file not found: nonexistent.log"

### Test 8: Invalid time format in filter
```bash
go run main.go -file=app_normal.log -from="invalid-date"
```
Expected error: "Error parsing timestamp: invalid-date"

### Test 9: CSV format (bonus)
```bash
go run main.go -file=app_normal.log -format=csv
```
Expected (if implemented): CSV lines like:
```
2024-01-15 10:30:45,INFO,Application started
2024-01-15 10:31:02,INFO,User login: admin
```

### Test 10: Keyword search (bonus)
```bash
go run main.go -file=app_large.log -keyword="Database"
```
Expected: Only entries containing "Database" in message

---

## Summary of Test Coverage

✓ Basic filtering (by level, time, both)
✓ Multiple output formats (text, JSON, CSV)
✓ Statistics calculation and accuracy
✓ Error handling (empty files, malformed entries)
✓ Edge cases (case insensitivity, special chars)
✓ Large file handling
✓ Boundary conditions (inclusive ranges)
✓ Missing flags (defaults work)
✓ Combining multiple filters
