package leetcode

/*
Problem statement

Given an arbitrary ransom note string and another string containing letters from all the magazines,
 write a function that will return true if the ransom note can be constructed from the magazines ;
  otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

*/

func canConstruct(ransomNote string, magazine string) bool {

	dict := map[rune]int{}

	for _, c := range magazine {
		dict[c]++
	}

	for _, c := range ransomNote {

		if dict[c] <= 0 {
			return false
		}
		if dict[c] > 0 {
			dict[c]--
		}
	}
	return true

}
