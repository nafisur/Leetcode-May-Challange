package leetcode

/*
Counting Bits
Solution
Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

Example 1:

Input: 2
Output: [0,1,1]
Example 2:

Input: 5
Output: [0,1,1,2,1,2]
*/
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}

	dp := make([]int, num+1)
	dp[0] = 0
	dp[1] = 1

	factor, temp := 2, 2
	for temp <= num {
		dp[temp] = 1
		temp *= 2
	}

	for i := 3; i <= num; i++ {
		if dp[i] != 0 {
			factor = i
		} else {
			dp[i] = dp[i-factor] + 1
		}
	}
	return dp
}
