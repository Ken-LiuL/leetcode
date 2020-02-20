package sort

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(num []int, l, r, target int) int {
	if r >= l {
		mid := (l + r) / 2
		if num[mid] > target {
			return binarySearch(num, l, mid-1, target)
		} else if num[mid] < target {
			return binarySearch(num, mid+1, r, target)
		} else {
			return mid
		}
	} else {
		return -1
	}
}