package solution

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		cnt1 := bits.OnesCount(uint(arr[i]))
		cnt2 := bits.OnesCount(uint(arr[j]))

		if cnt1 == cnt2 {
			return arr[i] < arr[j]
		}

		return cnt1 < cnt2
	})

	return arr
}
