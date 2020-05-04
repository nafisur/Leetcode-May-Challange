package leetcode

/*
Problem statement

Given a positive integer, output its complement number.
The complement strategy is to flip the bits of its binary representation.
Explanation: The binary representation of 5 is 101 (no leading zero bits),
and its complement is 010. So you need to output 2.
*/

import (
	"math"
)

func findComplement(num int) int {

	var result, i float64
	for num > 0 {
		if num%2 == 0 {
			result += math.Pow(2, i)
		}
		i++
		num /= 2
	}

	return int(result)
}
