package bitmanipulations

// isEven check if given no is even using bit manipulations
// we can do it by checking last bit of given no
// for Even last bit will be 1
// and for odd last bit will be 0
func isEven(num int) bool {
	// we can do & with mask `1` with given no
	if num&1 == 0 {
		return true
	}
	return false
}

func swap(a int, b int) []int {
	// we can do it by performing ^(XOR) on a and b
	a = a ^ b
	b = a ^ b // now b contains value of a
	a = a ^ b // now a will have value of b

	return []int{a, b}
}

// find i^th bit of given no
//  example - 5 = 101
// k = 2
// output = 0
func getKthBits(num int, k int) int {
	var mask = 1 << k
	if num&mask == 0 {
		return 0
	}
	return 1
}

func setKthBits(num int, k int) int {
	var mask = 1 << k
	return num | mask
}

func binaryToDecimal(binary []int) int {
	var decimal int
	for _, digit := range binary {
		decimal = decimal*2 + digit
	}

	return decimal
}
