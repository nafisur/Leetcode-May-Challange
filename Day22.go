package leetcode

/*
Sort Characters By Frequency
Solution
Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

Input:
"tree"

Output:
"eert"

Example 2:

Input:
"cccaaa"

Output:
"cccaaa"
*/
import "strings"

func frequencySort(s string) string {

	ls := len(s)
	freq := make([]int, 128)
	str := make([]string, ls+1)
	var result string

	for i := 0; i < ls; i++ {
		freq[s[i]]++
	}
	for i, v := range freq {
		if v != 0 {
			str[v] += strings.Repeat(string(i), v)
		}
	}
	for i := ls; i > 0; i-- {
		if str[i] != "" {
			result += str[i]
		}
	}
	return result
}
