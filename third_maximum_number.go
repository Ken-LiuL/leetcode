package array

import "math"

func thirdMax(nums []int) int {
	max, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for _, n := range nums {
		if n == max || n == second || n == third {
			continue
		}
		if n > max {
			third = second
			second = max
			max = n
		} else if n > second {
			third = second
			second = n
		} else if n > third {
			third = n
		}
	}
	if second == math.MinInt64 || third == math.MinInt64 {
		return  max
	}
	return third
}

