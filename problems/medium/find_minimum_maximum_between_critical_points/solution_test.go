package findminimummaximumbetweencriticalpoints

import (
	"reflect"
	"testing"
)

func TestNodesBetweenCriticalPoints(t *testing.T) {
	// Helper function to build lists cleanly
	buildList := func(vals []int) *ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &ListNode{Val: vals[0]}
		curr := head
		for _, v := range vals[1:] {
			curr.Next = &ListNode{Val: v}
			curr = curr.Next
		}
		return head
	}

	testCases := []struct {
		name     string
		vals     []int
		expected []int
	}{
		// {
		// 	name:     "Multiple critical points: [1,3,2,4,1]",
		// 	vals:     []int{1, 3, 2, 4, 1},
		// 	expected: []int{2, 2}, // Critical points at idx 2, 4 -> min: 2, max: 2
		// },
		// {
		// 	name:     "No critical points: [1,2,3,4]",
		// 	vals:     []int{1, 2, 3, 4},
		// 	expected: []int{-1, -1},
		// },
		// {
		// 	name:     "Only 1 critical point: [1,2,1]",
		// 	vals:     []int{1, 2, 1},
		// 	expected: []int{-1, -1}, // Needs at least 2 critical points
		// },
		// {
		// 	name:     "Example 2 from LeetCode: [5,3,1,2,5,1,2]",
		// 	vals:     []int{5, 3, 1, 2, 5, 1, 2},
		// 	expected: []int{1, 3},
		// },
		{
			name:     "Example 3: [1,3,2,2,3,2,2,2,7]",
			vals:     []int{1, 3, 2, 2, 3, 2, 2, 2, 7},
			expected: []int{3, 3}, // Critical points at idx 2 and 5 -> min: 3, max: 3
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := buildList(tc.vals)
			result := nodesBetweenCriticalPoints(head)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("nodesBetweenCriticalPoints(%v) = %v, want %v", tc.vals, result, tc.expected)
			}
		})
	}
}