package solution

// https://leetcode.com/problems/minimum-operations-to-equalize-binary-string
func minOperations(s string, k int) int {
	n := len(s)
	zeros := 0

	for _, char := range s {
		if char == '0' {
			zeros++
		}
	}

	if zeros == 0 {
		return 0
	}

	minZeros := zeros
	maxZeros := zeros

	for step := 1; step <= n; step++ {
		nextMin := 0
		if k < minZeros {
			nextMin = minZeros - k
		} else if k > maxZeros {
			nextMin = k - maxZeros
		} else if (k-minZeros)%2 != 0 {
			nextMin = 1
		} else {
			nextMin = 0
		}

		targetForMax := n - k
		nextMax := 0

		if targetForMax < minZeros {
			nextMax = n - (minZeros - targetForMax)
		} else if targetForMax > maxZeros {
			nextMax = n - (targetForMax - maxZeros)
		} else if (targetForMax-minZeros)%2 != 0 {
			nextMax = n - 1
		} else {
			nextMax = n
		}

		if nextMin == 0 {
			return step
		}

		minZeros = nextMin
		maxZeros = nextMax
	}

	return -1
}
