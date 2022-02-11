package slidingwindows

import "math"

func maxSumSubArrayOfSizeK(k int, nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	var maxSum = math.Inf(-1)
	var currentWindowSum = 0
	var left = 0
	for index := 0; index < len(nums)-k+2; index++ {
		currentWindowSum += nums[index]

		if index >= k {
			currentWindowSum -= nums[left]
			left++
		}

		maxSum = math.Max(float64(currentWindowSum), maxSum)
	}

	return int(maxSum)
}
