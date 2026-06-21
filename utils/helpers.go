package utils

// Add common helper functions here
// Examples:
// - Array utilities (sorting, searching, etc.)
// - String utilities
// - Tree/Graph helpers
// - Common data structures

func PrintArray(arr []int) {
	for _, val := range arr {
		print(val, " ")
	}
	println()
}

func PrintBoolArray(arr []bool) {
	for _, val := range arr {
		print(val, " ")
	}
	println()
}