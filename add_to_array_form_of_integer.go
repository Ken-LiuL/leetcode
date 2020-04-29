package array


func addToArrayForm(A []int, K int) []int {
	i, carry :=  len(A) - 1, 0
	for i >= 0 || K != 0 {
		sum := carry
		if i >= 0 {
			sum += A[i]
		}
		if K != 0 {
			sum += K % 10
			K = K / 10
		}
		carry = sum / 10
		remain := sum % 10
		if i >= 0 {
			A[i] = remain
            i--

		} else {
			A = append([]int{remain}, A...)
		}
	}
	if carry != 0 {
		A = append([]int{1}, A...)
	}
	return A
}