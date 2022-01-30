package dp

func canSum(target int, nums []int, memo map[int]bool) bool {
	if _, ok := memo[target]; ok {
		return memo[target]
	}

	if target == 0 {
		return true
	}

	if target < 0 {
		return false
	}

	for _, num := range nums {
		if canSum(target-num, nums, memo) {
			memo[target] = true

			return memo[target]
		}
	}

	memo[target] = false
	return memo[target]
}

func CanSum(target int, nums []int) bool {
	var memo = make(map[int]bool)

	return canSum(target, nums, memo)
}
