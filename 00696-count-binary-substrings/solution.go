package solution

// https://leetcode.com/problems/count-binary-substrings/
func countBinarySubstrings(s string) int {
	res, last, cur := 0, 0, 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			res += min(last, cur)
			last = cur
			cur = 1
		}
	}

	return res + min(last, cur)
}
