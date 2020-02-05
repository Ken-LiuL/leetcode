package sort


//this is dutch partitioning problem
//see links http://users.monash.edu/~lloyd/tildeAlgDS/Sort/Flag/
//red, blue and unknwn three pointers
//000111???222
//   r  u  b
// if nums[u] == red:
//			 swap(nums[u], nums[r])
//			 r++, u++ 
// if nums[u] == white:
//           u++    
// if nums[u] == blue:
//           swap(nums[u], nums[b])
//           u++, b--
func sortColors(nums []int)  {
	redPointer, unknownPointer, bluePointer := 0, 0, len(nums) - 1
	for unknownPointer <= bluePointer {
		if nums[unknownPointer] == 0 {
			 nums[redPointer], nums[unknownPointer] = nums[unknownPointer], nums[redPointer]
			 redPointer++
			 unknownPointer++
		} else if nums[unknownPointer] == 1 {
			 unknownPointer++
		} else {
			nums[unknownPointer], nums[bluePointer] = nums[bluePointer], nums[unknownPointer]
			bluePointer--
		}
	}
}