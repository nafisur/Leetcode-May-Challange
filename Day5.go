package leetcode

/*
Problem statement

Given a string, find the first non-repeating character in it and return it's index.
If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
*/

func firstUniqChar(s string) int {
	dict := map[rune]int{}

	for _, c := range s {
		dict[c]++
	}

	for i, c := range s {
		k, _ := dict[c]
		if k == 1 {
			return i
		}
	}

	return -1
}
