package sequentialdigits

import "testing"

func TestSequentialDigits(t *testing.T) {
	test := []struct {
		low      int
		high     int
		expected []int
	}{
		{low: 100, high: 300, expected: []int{123, 234}},
		{low: 1000, high: 13000, expected: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
	}

	for _, tc := range test {
		result := sequentialDigits(tc.low, tc.high)
		if len(result) != len(tc.expected) {
			t.Errorf("For low=%d and high=%d, expected length %d but got %d", tc.low, tc.high, len(tc.expected), len(result))
			continue
		}
		for i := range result {
			if result[i] != tc.expected[i] {
				t.Errorf("For low=%d and high=%d, expected %v but got %v", tc.low, tc.high, tc.expected, result)
				break
			}
		}
	}
}