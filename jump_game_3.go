package graph

func canReach(arr []int, start int) bool {
	if 0 <= start && start < len(arr) && arr[start] >= 0 {
		arr[start] = -arr[start]
		return arr[start] == 0 || canReach(arr, start+arr[start]) || canReach(arr, start-arr[start])
	}
	return false
}
