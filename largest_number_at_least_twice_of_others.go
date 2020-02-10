package array

func dominantIndex(nums []int) int {
	largest, second := 0, 0
	index := -1
	for ind, n := range nums {
		if n > largest {
			second = largest 
			largest = n
			index = ind
		} else if n > second {
			second = n
		}
	}
	if largest >= second * 2 {
		return  index
	} else {
		return -1
	}
}