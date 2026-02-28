package solution

// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers
func concatenatedBinary(n int) int {
	out := 0
	mod := 1000000007

	length := 0
	nextPower := 1

	for i := 1; i <= n; i++ {
		if i == nextPower {
			length++
			nextPower *= 2
		}

		for j := 0; j < length; j++ {
			out = (out * 2) % mod
		}

		out = (out + i) % mod
	}

	return out
}
