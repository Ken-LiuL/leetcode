package math

import "strconv"
func addStrings(num1 string, num2 string) string {
	res := ""
	i, j, carry := len(num1) -  1, len(num2) - 1, 0
	for i >= 0 || j >= 0 {
		 sum := carry
		 if i >= 0 {
             sum += int(num1[i] - '0')
             i--
		 }
		 if j >= 0 {
             sum += int(num2[j] - '0')
             j--
		 }
		 carry = sum / 10
		 res  = strconv.Itoa(sum % 10) + res 
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
	
}