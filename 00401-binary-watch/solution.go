package solution

import (
	"fmt"
	"math/bits"
)

// https://leetcode.com/problems/binary-watch
func readBinaryWatch(turnedOn int) []string {
	var res []string

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if bits.OnesCount(uint(h))+bits.OnesCount(uint(m)) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return res
}
