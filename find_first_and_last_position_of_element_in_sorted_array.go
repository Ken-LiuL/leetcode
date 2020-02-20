package array

func binarySearch(num []int, l, r, target int) int {
	if r >= l {
		mid := (l + r) / 2
		if num[mid] > target {
			for mid > 0 {
				mid--
				if num[mid] != num[mid+1] {
                    mid++
					break
				}
			}
			return binarySearch(num, l, mid-1, target)
		} else if num[mid] < target {
			for mid < len(num) - 1 {
				mid++
				if num[mid-1] != num[mid] {
                    mid--
					break
				}
			}
			return binarySearch(num, mid+1, r, target)
		} else {
			return mid
		}
	} else {
		return -1
	}
}

func searchRange(nums []int, target int) []int {
	ind := binarySearch(nums, 0, len(nums)-1, target)
	if ind == -1 {
        return  []int{-1, -1}
	}
	start, end := ind, ind
	for start > 0 &&  nums[start-1] == target {
			start -= 1
	}
	for end < len(nums) - 1 && nums[end+1] == target {
			end += 1
	}
	return []int{start, end}
}