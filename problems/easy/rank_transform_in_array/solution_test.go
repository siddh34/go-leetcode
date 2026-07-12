package ranktransforminarray

import "testing"

func TestArrayRankTransform(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{40, 10, 20, 30}, []int{4, 1, 2, 3}},
		{[]int{100, 100, 100}, []int{1, 1, 1}},
		{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := arrayRankTransform(tt.arr); !equal(got, tt.want) {
				t.Errorf("arrayRankTransform(%v) = %v; want %v", tt.arr, got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
