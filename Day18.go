package leetcode

/*
Permutation in String
Solution
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1.
In other words, one of the first string's permutations is the substring of the second string.

Example 1:

Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input:s1= "ab" s2 = "eidboaoo"
Output: False
*/

func checkInclusion(s1 string, s2 string) bool {

	ls1 := len(s1)
	ls2 := len(s2)
	if ls2 < ls1 {
		return false
	}
	if ls1 == 0 {
		return true
	}

	m1 := make([]int, 26)
	m2 := make([]int, 26)
	for i := range s1 {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}
	if equals(m1, m2) {
		return true
	}
	for i := ls1; i < ls2; i++ {
		m2[s2[i]-'a']++
		m2[s2[i-ls1]-'a']--
		if s2[i] != s2[i-ls1] && equals(m1, m2) {
			return true
		}
	}
	return false
}

func equals(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
