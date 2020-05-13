package leetcode

/*

Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

Note:

The length of num is less than 10002 and will be ≥ k.
The given num does not contain any leading zero.
Example 1:

Input: num = "1432219", k = 3
Output: "1219"

Example 2:

Input: num = "10200", k = 1
Output: "200"
*/

func removeKdigits(num string, k int) string {

	if len(num) == k {
		return "0"
	}

	var stack []byte
	stack = append(stack, num[0])

	for i := 1; i < len(num); i++ {
		for ; k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1]; k-- {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, num[i])
	}

	for ; k > 0 && len(stack) > 0; k-- {
		stack = stack[:len(stack)-1]
	}

	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
