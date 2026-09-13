package imageoverlap

import "testing"

func TestLargestOverlap(t *testing.T) {
	img1 := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	img2 := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}
	expected := 3
	result := largestOverlap(img1, img2)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}