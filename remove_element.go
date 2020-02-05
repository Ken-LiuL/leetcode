package array

func removeElement(nums []int, val int) int {
	 occur := 0
	 //two pointer
	 start, end :=  0, len(nums)-1
	 for start <= end {
		if nums[start] == val {
			nums[start], nums[end] = nums[end], nums[start]
			end--
			occur++
		} else {
			start++
		}
	 }
	 return len(nums) - occur
}