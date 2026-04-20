package regexmatch

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", "", true},
		{"a", ".", true},
		{"a", "a*", true},
	}

	for _, tt := range tests {
		t.Run(tt.s+"|"+tt.p, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
