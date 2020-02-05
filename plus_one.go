package list

func plusOne(digits []int) []int {
	carry := 1
    for j := len(digits) -1 ; j >= 0 ; j-- {
		 sum := digits[j] + carry
		 digits[j] = sum % 10
		 carry = sum / 10
	}
	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}