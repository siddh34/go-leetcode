package findgcd

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	for i := min; i >= 1; i-- {
		if min%i == 0 && max%i == 0 {
			return i
		}
	}

	return 1
}