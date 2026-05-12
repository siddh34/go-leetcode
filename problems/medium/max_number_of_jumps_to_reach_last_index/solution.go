package maxnumberofjumpstoreachlastindex

// https://leetcode.com/problems/maximum-number-of-jumps-to-reach-the-last-index/?envType=daily-question&envId=2026-05-10

func maximumJumps(nums []int, target int) int {
    dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = -1
		for j := 0; j < i; j++ {
			if dp[j] != -1 && nums[i] - nums[j] >= -target && nums[i] - nums[j] <= target {
				dp[i] = max(dp[i], dp[j] + 1)
			} else if j == 0 && nums[i] - nums[j] >= -target && nums[i] - nums[j] <= target {
				dp[i] = max(dp[i], 1)
			}
		}
	}

	return dp[len(nums) - 1]
}