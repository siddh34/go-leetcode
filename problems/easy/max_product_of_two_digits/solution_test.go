package maxproductoftwodigits

import "testing"

func TestMaxProduct(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{input: 123, expected: 6},
		{input: 456, expected: 30},
		{input: 987, expected: 72},
		{input: 1001, expected: 1},
	}
	
	for _, tc := range testCases {
		result := maxProduct(tc.input)
		if result != tc.expected {
			t.Errorf("For input %d, expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}