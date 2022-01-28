package dp

import "strings"

/**
how construct given string using string in given array
example:-
target: abcdf
arr : ["a", "abc", "cd", "df"] --> ["abc", "df"]
we can return any possible solutions
we can use word from given array as many time as we want.
**/

func howConstruct(target string, words []string) []string {
	if target == "" {
		return []string{}
	}

	var result []string
	for _, word := range words {
		if strings.Index(target, word) == 0 {
			var suffix = strings.TrimPrefix(target, word)
			var suffixWays = howConstruct(suffix, words)
			suffixWays = append(suffixWays, word)
			result = append(result, suffix)
		}
	}

	return result
}
