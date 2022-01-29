package dp

import "strings"

/**
can construct given string using string in given array
example:-
target: abcdf
arr : ["a", "abc", "cd", "df"] --> true
**/

func canConstruct(target string, words []string, memo map[string]bool) bool {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == "" {
		return true
	}

	for _, word := range words {
		if strings.Index(target, word) == 0 {
			if canConstruct(strings.TrimPrefix(target, word), words, memo) {
				memo[target] = true
				return memo[target]
			}
		}
	}

	memo[target] = false

	return memo[target]
}

func CanConstruct(target string, words []string) bool {
	var m = make(map[string]bool)
	return canConstruct(target, words, m)
}

/**
If m = len(target)
   n = len(words) then,
   TC: = O(M^n) and SC: O(m) without memoization
   After memmoization:
   TC : O(m*n) and sc : O(m)
**/
