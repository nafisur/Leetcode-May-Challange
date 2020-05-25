package leetcode

/*
Uncrossed Lines
We write the integers of A and B (in the order they are given) on two separate horizontal lines.

Now, we may draw connecting lines: a straight line connecting two numbers A[i] and B[j] such that:

A[i] == B[j];
The line we draw does not intersect any other connecting (non-horizontal) line.
Note that a connecting lines cannot intersect even at the endpoints: each number can only belong to one connecting line.

Return the maximum number of connecting lines we can draw in this way.

Example 1:

Input: A = [1,4,2], B = [1,2,4]
Output: 2

Example 2:

Input: A = [2,5,1,2,5], B = [10,5,2,1,5,2]
Output: 3
Example 3:

Input: A = [1,3,7,1,7,5], B = [1,9,2,5,1]
Output: 2
*/
func maxUncrossedLines(A []int, B []int) int {

	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	arr := make([][]int, len(A)+1)
	for i := range arr {
		arr[i] = make([]int, len(B)+1)
	}

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				arr[i][j] = arr[i-1][j-1] + 1
			} else {
				arr[i][j] = max(arr[i-1][j], arr[i][j-1])
			}
		}
	}
	return arr[len(A)][len(B)]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
