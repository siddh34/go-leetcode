package rotationfunction

import "testing"

func TestMaxRotateFunction(t *testing.T) {
	results := []struct {
		nums []int
		want int
	}{
		{[]int{4, 3, 2, 6}, 26},
		{[]int{1, 2, 3}, 8},
		{[]int{0, 0, 0}, 0},
		{[]int{1}, 0},
	}
	for _, tt := range results {
		t.Run("", func(t *testing.T) {
			if got := maxRotateFunction(tt.nums); got != tt.want {
				t.Errorf("maxRotateFunction(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
