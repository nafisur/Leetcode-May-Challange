package leetcode

/*
Find All Anagrams in a String
Solution
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:
Input:
s: "cbaebabacd" p: "abc"
Output:
[0, 6]

Example 2:
Input:
s: "abab" p: "ab"
Output:
[0, 1, 2]
*/

func findAnagrams(s string, p string) []int {

	ls := len(s)
	lp := len(p)
	if ls < lp {
		return nil
	}
	var (
		m1     [26]int
		m2     [26]int
		result []int
	)
	for i := range p {
		m1[p[i]-'a']++
		m2[s[i]-'a']++
	}
	for i := 0; i < ls-lp+1; i++ {
		if m1 == m2 {
			result = append(result, i)
		}
		if i+lp < ls {
			m2[s[i]-'a']--
			m2[s[i+lp]-'a']++
		}
	}
	return result
}
