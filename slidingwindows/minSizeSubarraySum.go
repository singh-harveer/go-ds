package slidingwindows

import "math"

/**
Problem Statement#
Given an array of positive numbers and a positive number ‘S,’ find the length of the smallest contiguous
subarray whose sum is greater than or equal to ‘S’. Return 0 if no such subarray exists.
example -1
Input: [2, 1, 5, 2, 3, 2], S=7
Output: 2
Explanation: The smallest subarray with a sum greater than or equal to '7' is [5, 2].

Example -2
Input: [2, 1, 5, 2, 8], S=7
Output: 1
Explanation: The smallest subarray with a sum greater than or equal to '7' is [8].
**/

func minSizeSubArraySum(nums []int, sum int) int {
	var currentSum = 0
	var left = 0
	var subArrayMinSize = math.Inf(1)

	for index, num := range nums {
		if num >= sum {
			subArrayMinSize = 1
			continue
		}

		currentSum += num

		if currentSum >= sum {
			currentSum -= nums[left] // sliding window

			var currentSubArraySize = index - left
			subArrayMinSize = math.Min(float64(currentSubArraySize), subArrayMinSize)
			left++
		}
	}

	if math.IsInf(subArrayMinSize, 1) {
		return 0
	}

	return int(subArrayMinSize)
}
