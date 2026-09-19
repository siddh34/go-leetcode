package circleandrectangleoverlapping

import "testing"

func TestCheckOverlap(t *testing.T) {
	test := []struct {
		radius   int
		xCenter  int
		yCenter  int
		x1       int
		y1       int
		x2       int
		y2       int
		expected bool
	}{
		{1, 0, 0, -1, -1, 1, 1, true},
		{1, 0, 0, 2, 2, 3, 3, false},
	}

	for _, tt := range test {
		result := checkOverlap(tt.radius, tt.xCenter, tt.yCenter, tt.x1, tt.y1, tt.x2, tt.y2)
		if result != tt.expected {
			t.Errorf("checkOverlap(%d, %d, %d, %d, %d, %d, %d) = %v; want %v", tt.radius, tt.xCenter, tt.yCenter, tt.x1, tt.y1, tt.x2, tt.y2, result, tt.expected)
		}
	}
}