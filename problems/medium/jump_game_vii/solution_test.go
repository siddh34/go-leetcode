package jumpgamevii

import "testing"

func TestCanReach(t *testing.T) {
	type args struct {
		s       string
		minJump int
		maxJump int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s:       "011010",
				minJump: 2,
				maxJump: 3,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s:       "01101110",
				minJump: 2,
				maxJump: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReach(tt.args.s, tt.args.minJump, tt.args.maxJump); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
