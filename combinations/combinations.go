package combinations

/**
find combinations of given slice of elements
Example -1
input = [A,B,C]
output = [[],
			[A],[B], [C],
			[A,B],[B, C], [C,A],
			[A,B,C]
		]
Example -2
input = [] // base case
output = [ [] ]
**/

func combinationOfChar(arr []string) [][]string {
	if len(arr) == 0 {
		return [][]string{{}}
	}

	var first = arr[0]
	var rest = arr[1:]
	var combsWithoutFirst = combinationOfChar(rest)

	var combsWithFirst [][]string

	for _, comb := range combsWithoutFirst {
		comb = append(comb, first)
		combsWithoutFirst = append(combsWithoutFirst, comb)
	}

	return append(combsWithFirst, combsWithoutFirst...)
}
