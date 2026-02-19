package solution

// https://leetcode.com/problems/compare-version-numbers/
func compareVersion(version1 string, version2 string) int {
	n1, n2 := len(version1), len(version2)
	i, j := 0, 0

	for i < n1 || j < n2 {
		var num1, num2 int

		for i < n1 && version1[i] != '.' {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}

		for j < n2 && version2[j] != '.' {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}

		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}

		i++
		j++
	}

	return 0
}
