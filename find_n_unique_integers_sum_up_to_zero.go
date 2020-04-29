package array

func sumZero(n int) []int {
	res := make([]int, n)
	//odd
	if n%2 == 1 {
		start := -(n / 2)
		for i := 0; i < n; i++ {
			res[i] = start
			start++
		}
	} else {
		start := -(n / 2)
		for i := 0; i < n; {
			if start != 0 {
				res[i] = start
				i++
			}
			start++
		}
	}
	return res
}
