package array


func missingNumber(nums []int) int {
	l  := len(nums)
	sum := l * (l + 1) / 2

	for _, n := range nums {
		sum -= n
	}
	return sum
}