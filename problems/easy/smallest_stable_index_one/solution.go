package smallest_stable_index_one

import (
	"fmt"
	"math"
)

func firstStableIndex(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}

	minStableIndex := -1
    for i := 0; i < len(nums); i++ {
		maxNumber := 0
		minNumber := math.MaxInt64

		for k := 0; k <= i; k++ {
			if nums[k] > maxNumber {
				maxNumber = nums[k]
			}
		}

		for j := i; j < len(nums); j++ {
			if nums[j] < minNumber {
				minNumber = nums[j]
			}
		}

		stability := maxNumber - minNumber
		if stability <= k && minStableIndex < i {
			if minStableIndex == -1 {
				minStableIndex = i
			} else {
				minStableIndex = min(minStableIndex, i)
			}
		}
		fmt.Println("i:", i, "maxNumber:", maxNumber, "minNumber:", minNumber, "minStableIndex:", minStableIndex, "stability:", stability)
	}

	fmt.Println("minStableIndex:", minStableIndex)
	return minStableIndex
}
