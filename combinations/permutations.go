package combinations

/**
find permutations of given slice of String Char
Example -1
input = [A,B,C]
output = [][]string{
		{"A", "B", "C"},
		{"A", "C", "B"},
		{"C", "A", "B"},
		{"C", "B", "A"},
		{"B", "C", "A"},
		{"B", "A", "C"}
		}
Example -2
input = [] // base case
output = [][]string{{}}
**/

func permutationOfChar(arr []string) [][]string {
	if len(arr) == 0 {
		return [][]string{{}}
	}

	var first, rest = arr[0], arr[1:]
	var permsWithoutFirst = permutationOfChar(rest)

	var allPermutations [][]string

	for _, perm := range permsWithoutFirst {
		for index := 0; index <= len(perm); index++ {
			// insert 'first' at every index
			var permsWithFirst = append(perm[:index], append([]string{first}, perm[index:]...)...)
			allPermutations = append(allPermutations, permsWithFirst)
		}
	}

	return allPermutations
}
