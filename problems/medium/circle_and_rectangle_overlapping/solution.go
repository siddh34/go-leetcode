package circleandrectangleoverlapping

import "math"

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    dx := math.Max(float64(x1-xCenter), math.Max(0, float64(xCenter-x2)))
    dy := math.Max(float64(y1-yCenter), math.Max(0, float64(yCenter-y2)))
    return dx*dx+dy*dy <= float64(radius*radius)
}