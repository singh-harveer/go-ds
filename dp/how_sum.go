package dp

func howSum(target int, nums []int, memo map[int][]int) []int {
	if _, ok := memo[target]; ok {
		return memo[target]
	}

	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	var result []int
	for _, num := range nums {
		var waysToSum = howSum(target-num, nums, memo)
		if waysToSum != nil {
			result = append(result, num)
			result = append(result, waysToSum...)
			memo[target] = result

			return memo[target]
		}
	}
	memo[target] = result

	return memo[target]
}

func HowSum(target int, nums []int) []int {
	var memo = make(map[int][]int)

	return howSum(target, nums, memo)
}
