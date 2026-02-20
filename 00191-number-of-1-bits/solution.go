package solution

import "math/bits"

// https://leetcode.com/problems/number-of-1-bits
func hammingWeight(n int) int {
	return bits.OnesCount(uint(n))
}
