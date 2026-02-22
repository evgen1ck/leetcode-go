package solution

// https://leetcode.com/problems/binary-gap
func binaryGap(n int) int {
	maxGap := 0
	lastPos := -1

	for i := 0; n > 0; i++ {
		if n%2 == 1 {
			if lastPos != -1 && i-lastPos > maxGap {
				maxGap = i - lastPos
			}
			lastPos = i
		}

		n /= 2
	}

	return maxGap
}
