package array
import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return helper(candidates, target , 0)
}

func helper(candidates []int, target int, start int) [][]int {
	res := make([][]int, 0, 10)
	for i := start; i < len(candidates); i++ {
		 c := candidates[i]
		 if target < c {
			 break 
		 } 
		 if target - c == 0 {
			 res = append(res, []int{c})
		 } else {
		   r := helper(candidates, target - c, i) 
		   for i := 0; i < len(r); i++ {
			 r[i] = append(r[i], c)
			 res = append(res, r[i])
		   }
		}
	}
	return res
}