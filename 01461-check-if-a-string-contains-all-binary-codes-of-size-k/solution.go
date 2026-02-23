package solution

import "math"

// https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k
func hasAllCodes(s string, k int) bool {
	targetCount := int(math.Pow(2, float64(k)))
	seen := make(map[string]bool)

	for i := 0; i <= len(s)-k; i++ {
		substring := s[i : i+k]
		seen[substring] = true
	}

	return len(seen) == targetCount
}
