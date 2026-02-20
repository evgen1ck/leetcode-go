package solution

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/special-binary-string
func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}

	var parts []string
	cnt, left := 0, 0

	for i, char := range s {
		if char == '1' {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			inner := makeLargestSpecial(s[left+1 : i])
			parts = append(parts, "1"+inner+"0")
			left = i + 1
		}
	}

	sort.Slice(parts, func(i, j int) bool {
		return parts[i] > parts[j]
	})

	return strings.Join(parts, "")
}
