package list

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, 100000)
	stack := make([]int, 0, len(nums2))
	for _, n := range nums2 {
	   for len(stack) != 0 && n > stack[len(stack) - 1] {
		   res[stack[len(stack) - 1]] = n
		   stack = stack[:len(stack) - 1]
	   }
	   stack = append(stack, n)
	}
	for _, i := range stack {
		res[i] = -1
	}
	ret := make([]int, len(nums1))
	for i, n := range nums1 {
		ret[i] = res[n]
	}

	return ret

}