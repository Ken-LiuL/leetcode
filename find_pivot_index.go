package array


func pivotIndex(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	res :=  0
	for i := 0; i < len(nums); i++ {
		if res * 2 == sum - nums[i] {
			return i
		}
		res += nums[i]
	}
	return -1
}