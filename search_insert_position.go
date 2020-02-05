package array


func searchInsert(nums []int, target int) int {
	 start, end := 0, len(nums)-1
	 
	 for start <= end {
		 mid := start + (end - start)/2
		 if nums[mid] < target {
			 start = mid + 1 
		 } else if nums[mid] == target {
			 return mid
		 } else {
			 end = mid - 1
		 }
	 }
	 return  start
}