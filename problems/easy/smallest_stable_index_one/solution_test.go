package smallest_stable_index_one

import "testing"

func TestFirstStableIndex(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{5,0,1,4},
			k:    3,
			want: 3,
		},
		{
			nums: []int{3,2,1},
			k:    1,
			want: -1,
		},
		{
			nums: []int{0, 0},
			k:    0,
			want: 0,
		},
	}
	for _, tt := range tests {
		got := firstStableIndex(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("firstStableIndex(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}