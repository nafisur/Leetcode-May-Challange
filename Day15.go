package leetcode

/*
Maximum Sum Circular Subarray
Solution
Given a circular array C of integers represented by A, find the maximum possible sum of a non-empty subarray of C.

Here, a circular array means the end of the array connects to the beginning of the array.  (Formally, C[i] = A[i] when 0 <= i < A.length, and C[i+A.length] = C[i] when i >= 0.)

Also, a subarray may only include each element of the fixed buffer A at most once.  (Formally, for a subarray C[i], C[i+1], ..., C[j], there does not exist i <= k1, k2 <= j with k1 % A.length = k2 % A.length.)
Example 1:

Input: [1,-2,3,-2]
Output: 3
Explanation: Subarray [3] has maximum sum 3
Example 2:

Input: [5,-3,5]
Output: 10
Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10
Example 3:

Input: [3,-1,2,-1]
Output: 4
Explanation: Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4
*/

import "math"

func maxSubarraySumCircular(A []int) int {

	if len(A) <= 1 {
		return A[0]
	}
	for i := 0; i < len(A); i++ {
		A[i] = -A[i]
	}
	modifiedMax := kadanes(A)

	for i := 0; i < len(A); i++ {
		A[i] = -A[i]
	}
	sumOfA := sum(A)
	kad := kadanes(A)
	if sumOfA+modifiedMax <= 0 {
		return kad
	}
	return max(kadanes(A), sumOfA+modifiedMax)
}

func kadanes(A []int) int {

	maximum := int(math.MinInt32)
	temp := A[0]

	for i := 1; i < len(A); i++ {
		temp = max(A[i], temp+A[i])
		maximum = max(temp, maximum)
	}
	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(input []int) int {
	var sum int
	for _, v := range input {
		sum += v
	}
	return sum
}
