package maxnumberofjumpstoreachlastindex

import "testing"

func TestMaxJumps(t *testing.T) {
	results := []struct {
		arr  []int
		tgt  int
		want int
	}{
		{[]int{1,3,6,4,1,2}, 2, 3},
		{[]int{1,3,6,4,1,2}, 3, 5},
		{[]int{1,3,6,4,1,2}, 0, -1},
	}
	for _, tt := range results {
		t.Run("", func(t *testing.T) {
			if got := maximumJumps(tt.arr, tt.tgt); got != tt.want {
				t.Errorf("maximumJumps(%v, %d) = %d, want %d", tt.arr, tt.tgt, got, tt.want)
			}
		})
	}
}