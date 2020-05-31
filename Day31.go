package leetcode

/*
Edit Distance
Solution
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	cost := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		cost[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		cost[i][0] = i
	}
	for i := 1; i <= n; i++ {
		cost[0][i] = i
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				cost[i+1][j+1] = cost[i][j]
			} else {
				cost[i+1][j+1] = 1 + min(cost[i][j], min(cost[i][j+1], cost[i+1][j]))
			}
		}
	}

	return cost[m][n]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
