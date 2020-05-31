package leetcode

/*
Possible Bipartition
Solution
Given a set of N people (numbered 1, 2, ..., N), we would like to split everyone into two groups of any size.

Each person may dislike some other people, and they should not go into the same group.

Formally, if dislikes[i] = [a, b], it means it is not allowed to put the people numbered a and b into the same group.

Return true if and only if it is possible to split everyone into two groups in this way.

Example 1:

Input: N = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: group1 [1,4], group2 [2,3]
Example 2:

Input: N = 3, dislikes = [[1,2],[1,3],[2,3]]
Output: false
Example 3:

Input: N = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
Output: false
*/
var color map[int]int
var edges map[int][]int

func possibleBipartition(N int, dislikes [][]int) bool {
	color = make(map[int]int)
	edges = make(map[int][]int)

	for _, dislike := range dislikes {
		edges[dislike[0]] = append(edges[dislike[0]], dislike[1])
		edges[dislike[1]] = append(edges[dislike[1]], dislike[0])
	}

	for i := 1; i <= N; i++ {

		if _, ok := color[i]; !ok {
			if !search(i, 0) {
				return false
			}
		}
	}

	return true

}

func search(node, c int) bool {

	if val, ok := color[node]; ok {
		return val == c
	}
	color[node] = c

	for _, n := range edges[node] {

		if !search(n, c^1) {
			return false
		}
	}

	return true

}
