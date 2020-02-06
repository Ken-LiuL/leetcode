package array

//https://gregable.com/2013/10/majority-vote-algorithm-find-majority.html
func majorityElement(nums []int) int {
	 count, res := 0, nums[0]
	 for _,  n := range nums[1:] {
		 if n == res {
			count++
		 } else if count == 0 {
			 res = n
		 } else {
			 count--
		 }
	 }
	 return res

}