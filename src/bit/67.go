package bit

import "strconv"

// 67. 二进制求和
// 给你两个二进制字符串，返回它们的和（用二进制表示）。


func addBinary(a string, b string) string {
	carry := 0
	result := ""

	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		carry = sum / 2
		result = strconv.Itoa(sum%2) + result
	}
	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}
	return result
}

