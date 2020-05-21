package leetcode

/*
Problem statement

Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.
Example 1:

Input: 16
Output: true


Example 2:

Input: 14
Output: false

*/

func isPerfectSquare(num int) bool {

	if num < 2 {
		return true
	}

	start := 1
	end := num
	for start <= end {
		mid := (start + end) / 2
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}
	return false
}
