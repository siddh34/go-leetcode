package smallestsubsequenceofcharacters

import "testing"

func TestSmallestSubsequence(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "bcabc", expected: "abc"},
		{input: "cbacdcbc", expected: "acdb"},
	}

	for _, tc := range testCases {
		result := smallestSubsequence(tc.input)
		if result != tc.expected {
			t.Errorf("For input %s, expected %s but got %s", tc.input, tc.expected, result)
		}
	}
}