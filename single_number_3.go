package array

func singleNumber(nums []int) []int {
	res := make([]int, 2)
	// get xor res
	xorRes := nums[0]  
	for _, n := range nums[1:] {
		 xorRes ^= n
	}
	//get mask: get the lowest one bit
	mask := xorRes & -xorRes
	//split two group
	for _, n := range nums {
		if n & mask == 0 {
			res[0] ^= n
		} else {
			res[1] ^= n
		}
	}
	return res
}