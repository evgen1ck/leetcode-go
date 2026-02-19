package solution

// https://leetcode.com/problems/binary-number-with-alternating-bits/
func hasAlternatingBits(n int) bool {
	n = n ^ (n >> 1)
	return (n & (n + 1)) == 0
}
