package arrays

/**
Given an array of integers, check if there is a pair which is equal to provided sum
array = [1,2,9,4]
sum = 10

Output: 1, 9


array = [1,2,5,4]
sum = 8
**/

func pairSum(target int, nums []int) []int {
	var m = make(map[int]struct{})

	var result []int
	for _, num := range nums {
		var key = target - num

		if _, ok := m[key]; ok {
			result = append(result, key)
			result = append(result, num)

			return result
		}
		m[num] = struct{}{}
	}

	return nil
}
