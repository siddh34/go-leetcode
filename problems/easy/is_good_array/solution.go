package isgoodarray

func isGood(nums []int) bool {
	max := 0
	repeat := make(map[int]int)

	for _, num := range nums {
		if num > max {
			max = num
		}
		repeat[num]++
	}

	if repeat[max] != 2 {
		return false
	}

	for i := 1; i < max; i++ {
		if repeat[i] != 1 {
			return false
		}
	}

	return true
}