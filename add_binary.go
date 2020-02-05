package str

import "strconv"


func addBinary(a string, b string) string {
	res := ""
	i, j, carry := len(a) -  1, len(b) - 1, 0
	for i >= 0 || j >= 0 {
		sum := carry
		if j >= 0 {
		  sum += int(b[j] -'0')
		  j--
		}
		if i >= 0 {
		  sum += int(a[i] - '0')
		  i--
		}
		res = strconv.Itoa(sum % 2) + res
		carry = sum / 2
	}
	if carry != 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}