package isgoodarray

import "testing"

func TestIsGood(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "test1",
			nums: []int{1, 2, 3, 3},
			want: true,
		},
		{
			name: "test2",
			nums: []int{1, 1, 2, 3},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGood(tt.nums); got != tt.want {
				t.Errorf("isGood() = %v, want %v", got, tt.want)
			}
		})
	}
}