package slidingwindows

import "sort"

func AllTriplets(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for index, num := range nums {
		// to handle duplicate num in slice
		if index != 0 && num == nums[index-1] {
			continue
		}

		var currentTriplets = pairSum(nums, -num, index+1)
		result = append(result, currentTriplets...)
	}

	return result
}

func pairSum(nums []int, targetSum int, left int) [][]int {
	var right = len(nums) - 1

	var triplets [][]int
	for left < right {
		if nums[left]+nums[right] == targetSum {
			triplets = append(triplets, []int{-targetSum, nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}

			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if nums[left]+nums[right] > targetSum {
			right--
		} else {
			left++
		}
	}

	return triplets
}
