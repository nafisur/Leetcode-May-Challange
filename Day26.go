package leetcode

/*
Contiguous Array
Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

Example 1:
Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
Example 2:
Input: [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
Note: The length of the given binary array will not exceed 50,000.

*/
func findMaxLength(nums []int) int {
	m := map[int]int{} //key is count, value is index

	var count, maxVal int
	m[0] = -1
	for i, num := range nums {
		if num == 0 {
			count--
		} else {
			count++
		}
		idx, ok := m[count]
		if ok {
			maxVal = max(maxVal, i-idx)
		} else {
			m[count] = i
		}
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
