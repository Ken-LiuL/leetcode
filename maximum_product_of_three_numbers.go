package array

func max(a, b int) int {
	if a > b  {
		return a
	}
	return b
}


func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0]*nums[1]*nums[2]
	}
	min1, min2, max1, max2, max3 := 1000, 1000, -1000, -1000, -1000
	for _, n := range nums {
		if n < min1 {
			min2 = min1
			min1 = n
		} else if n < min2 {
			min2 = n
		} 
		if n > max1  {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 {
			max3 = max2
			max2 = n
		} else if n > max3 {
			max3 = n 
		}
	} 
	return max(min1*min2*max1, max1*max2*max3)
}