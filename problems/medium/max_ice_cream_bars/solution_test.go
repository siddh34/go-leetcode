package maxicecreambars

import "testing"

func TestMaxIceCream(t *testing.T) {
	type args struct {
		costs []int
		coins int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				costs: []int{1, 3, 2, 4, 1},
				coins: 7,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				costs: []int{10, 6, 8, 7, 7, 8},
				coins: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIceCream(tt.args.costs, tt.args.coins); got != tt.want {
				t.Errorf("maxIceCream() = %v, want %v", got, tt.want)
			}
		})
	}
}