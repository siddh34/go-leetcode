package numberofuniquexortriplets

import "testing"

func TestUniqueXorTriplets(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{1,2}, expected: 2},
		{nums: []int{5, 1, 6}, expected: 4},
		{nums: []int{1, 1, 1}, expected: 4},
	}

	for _, tc := range testCases {
		result := uniqueXorTriplets(tc.nums)
		if result != tc.expected {
			t.Errorf("For nums %v, expected %d but got %d", tc.nums, tc.expected, result)
		}
	}
}
