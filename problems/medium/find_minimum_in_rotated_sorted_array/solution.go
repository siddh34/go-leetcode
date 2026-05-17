package findminimuminrotatedsortedarray

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/?envType=daily-question&envId=2026-05-15

func findMin(nums []int) int {
	left, right := 0, len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	
	return nums[left]
}