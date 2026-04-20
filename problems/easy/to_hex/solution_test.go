package tohex

import "testing"

func TestToHex(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{26, "1a"},
		{-1, "ffffffff"},
		{0, "0"},
	}

	for _, tt := range tests {
		t.Run(string(rune(tt.num)), func(t *testing.T) {
			if got := toHex(tt.num); got != tt.want {
				t.Errorf("toHex(%d) = %q, want %q", tt.num, got, tt.want)
			}
		})
	}
}