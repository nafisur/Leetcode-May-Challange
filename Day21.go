package leetcode

/*
Count Square Submatrices with All Ones
Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
Input: matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
Output: 15
Explanation:
There are 10 squares of side 1.
There are 4 squares of side 2.
There is  1 square of side 3.
Total number of squares = 10 + 4 + 1 = 15.
*/
func countSquares(matrix [][]int) int {

	var result int

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					result++
					continue
				}
				matrix[i][j] = 1 + min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1]))
				result += matrix[i][j]
			}
		}
	}
	return result

}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
