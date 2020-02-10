package array


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
//1, 2, 6, 4, 8, 10, 9, 15 20
//find the right pivot, all elements to the right of the pivot is larger than left side
//find the left pivot, all elements to the left of the pivot is  smaller than right side
func findUnsortedSubarray(nums []int) int {
	start, end :=  -1, -2
	maxV, minV := nums[0], nums[len(nums) - 1]
	for i := 0; i < len(nums); i++ {
		if nums[i] < maxV {
			end = i
		}
		if nums[len(nums)-i-1] > minV {
			start = len(nums) - i -1
		}
		maxV =  max(nums[i], maxV)
		minV =  min(nums[len(nums) -i -1], minV)
	}
	return end - start + 1 
}