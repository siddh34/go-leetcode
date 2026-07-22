package shiftgrid

import "testing"

func TestShiftGrid(t *testing.T) {
	testCases := []struct {
		grid     [][]int
		k        int
		expected [][]int
	}{
		{
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        1,
			expected: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			grid:     [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			k:        4,
			expected: [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			grid:     [][]int{{1, 2}, {3, 4}, {5, 6}},
			k:        3,
			expected: [][]int{{4, 5}, {6, 1}, {2, 3}},
		},
	}

	for _, tc := range testCases {
		result := shiftGrid(tc.grid, tc.k)
		if !equal(result, tc.expected) {
			t.Errorf("For grid %v and k %d, expected %v but got %v", tc.grid, tc.k, tc.expected, result)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
