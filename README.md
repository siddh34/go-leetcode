# LeetCode Solutions in Go

A collection of LeetCode problem solutions organized by difficulty.

## Structure

- `problems/easy/` - Easy difficulty problems
- `problems/medium/` - Medium difficulty problems
- `problems/hard/` - Hard difficulty problems
- `utils/` - Shared utility functions

## How to Add a Problem

1. Create a new folder under the appropriate difficulty: `problems/{difficulty}/{problem_name}/`
2. Create `solution.go` with your solution
3. Create `solution_test.go` with test cases

Example:
```
problems/medium/regular_expression_match/
├── solution.go
└── solution_test.go
```

## Running Tests

```bash
go test ./...
```

## Problem Tracking

| # | Title | Difficulty | Status |
|---|-------|-----------|--------|
| 10 | Regular Expression Matching | Medium | In Progress |
