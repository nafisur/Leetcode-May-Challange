package leetcode

/*
Problem statement

You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.
Find this single element that appears only once.

Example 1:

Input: [1,1,2,3,3,4,4,8,8]
Output: 2
Example 2:

Input: [3,3,7,7,10,11,11]
Output: 10

Note: Your solution should run in O(log n) time and O(1) space.
*/

func singleNonDuplicate(nums []int) int {

	var (
		start int = 0
		end   int = len(nums) - 1
	)

	for start <= end {

		mid := start + (end-start)/2
		if start == end {
			return nums[mid]
		} else if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				start = mid + 2
			} else {
				end = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

	} //for
	return 0
}
