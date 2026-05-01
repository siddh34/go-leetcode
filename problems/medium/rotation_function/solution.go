package rotationfunction

// https://leetcode.com/problems/rotate-function/?envType=daily-question&envId=2026-05-01

/*
Using Recurrance formula

Observed as F(k) = F(k-1) + sum - n * nums[n-k] 
*/

func maxRotateFunction(nums []int) int {
	sum := 0
	n := len(nums)

	for _, num := range nums {
		sum += num
	}

	currentNumberInSeries := 0
	for i, num := range nums {
		currentNumberInSeries += i * num
	}

	maxNumberInSeries := currentNumberInSeries
	for k := 1; k < n; k++ {
		numberMoved := nums[n-k]
		currentNumberInSeries = currentNumberInSeries + sum - n*numberMoved
		maxNumberInSeries = max(currentNumberInSeries, maxNumberInSeries)
	}
	return maxNumberInSeries
}
