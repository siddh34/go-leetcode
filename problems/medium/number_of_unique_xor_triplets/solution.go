package numberofuniquexortriplets

import "math/bits"

func uniqueXorTriplets(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	uniqueXors := 1 << bits.Len(uint(max(n)))
	return uniqueXors
}