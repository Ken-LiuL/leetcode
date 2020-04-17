package list

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([][]int, 0, len(nums))
	res := make([]int,0, len(nums)) 
	for i, n := range nums {
		if i >=k && len(queue) != 0 && queue[0][0] + k <= i {
			queue = queue[1:]
		}
		for len(queue) != 0 && queue[len(queue)-1][1] < n {
		   queue = queue[:len(queue)-1]
		}
		queue = append(queue, []int{i, n})
		if i >= k-1 {
			res = append(res, queue[0][1])
		}
	}
	return res
}