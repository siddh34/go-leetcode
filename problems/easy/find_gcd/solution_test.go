package findgcd

import "testing"

func TestFindGCD(t *testing.T) {
	tests := []struct {
		nums    []int
		expected int
	} {
		{nums: []int{2, 5, 6, 9, 10}, expected: 2},
		{nums: []int{7, 5, 6, 8, 3}, expected: 1},
		{nums: []int{3, 3}, expected: 3},
	}
	for _, test := range tests {
		result := findGCD(test.nums)
		if result != test.expected {
			t.Errorf("For nums %v, expected %d, but got %d", test.nums, test.expected, result)
		}
	}
}
