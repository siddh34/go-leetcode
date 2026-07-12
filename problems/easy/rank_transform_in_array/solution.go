package ranktransforminarray

import (
	"slices"
)

func arrayRankTransform(arr []int) []int {
    copied_array := make([]int, len(arr))
	copy(copied_array, arr)
	slices.Sort(copied_array)

	rank_map := make(map[int]int)
	rank := 1
	for _, num := range copied_array {
		if _, exists := rank_map[num]; !exists {
			rank_map[num] = rank
			rank++
		}
	}

	result := make([]int, len(arr))
	for i, num := range arr {
		result[i] = rank_map[num]
	}
	return result
}
