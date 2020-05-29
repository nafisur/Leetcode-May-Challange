package leetcode

/*
Course Schedule
Solution
There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 {
		return false
	}
	if prerequisites == nil || len(prerequisites) == 0 {
		return true
	}
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		indegree[p[0]]++
	}

	var curr []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			curr = append(curr, i)
		}
	}

	for len(curr) != 0 {
		next := []int{}
		for _, course := range curr {
			for _, neighbor := range graph[course] {
				indegree[neighbor]--
				if indegree[neighbor] == 0 {
					next = append(next, neighbor)
				}
			}
		}
		curr = next
	}
	for i := 0; i < numCourses; i++ {
		if indegree[i] != 0 {
			return false
		}
	}
	return true
}
