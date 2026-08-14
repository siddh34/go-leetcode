package maximumlengthsubstringwithtwooccurrences

import "testing"

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"abca", 4},
		{"cbzxy", 5},
		{"abcabc", 6},
	}

	for _, test := range tests {
		result := maxLengthBetweenEqualCharacters(test.s)
		if result != test.expected {
			t.Errorf("For input %s, expected %d but got %d", test.s, test.expected, result)
		}
	}
}