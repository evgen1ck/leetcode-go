package solution

import (
	"math/big"
)

// https://leetcode.com/problems/add-binary
func addBinary(a string, b string) string {
	num1 := new(big.Int)
	num2 := new(big.Int)

	num1.SetString(a, 2)
	num2.SetString(b, 2)

	num1.Add(num1, num2)

	return num1.Text(2)
}
