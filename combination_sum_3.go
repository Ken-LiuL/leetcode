package array


func combinationSum3(k int, n int) [][]int {
	return helper(k , n , []int{1, 2, 3, 4 ,5,6, 7, 8, 9})
}

func helper(k int, n int, arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}
	res := make([][]int , 0, 10)
	for ind, a := range arr {
		if  a > n {
			break
		}
		if a == n && k == 1 {
            res = append(res, []int{a})
		} else if a == n {
			break
		} else if k == 1 {
			continue
		} else {
			for _ , r := range helper(k - 1, n - a, arr[ind+1: ]) {
				r = append(r, a)
				res = append(res, r)
			}
		}
	}
	return res
}