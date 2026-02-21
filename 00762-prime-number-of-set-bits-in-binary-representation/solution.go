package solution

import "math/bits"

// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation
func countPrimeSetBits(left int, right int) int {
	res := 0

	for i := left; i <= right; i++ {
		cnt := bits.OnesCount(uint(i))

		switch cnt {
		case 2, 3, 5, 7, 11, 13, 17, 19:
			res++
		}
	}

	return res
}
