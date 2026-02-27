package solution

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one
func numSteps(s string) int {
	steps := 0
	carry := 0

	for i := len(s) - 1; i > 0; i-- {
		bit := int(s[i]-'0') + carry

		if bit == 1 {
			steps += 2
			carry = 1
		} else {
			steps += 1
		}
	}

	if carry == 1 {
		steps += 1
	}

	return steps
}
