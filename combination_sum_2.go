package array

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return helper(candidates, target)
}
func helper(candidates []int, target int) [][]int {
	 if len(candidates) ==  0 {
		 return [][]int{}
	 } 
	 var res [][]int = make([][]int, 0, 10)
	 for i := 0; i < len(candidates); {
         if i > 0 && candidates[i] == candidates[i-1] {
			 i++ 
             continue
		 }
		 c := candidates[i]
		 if c > target {
			 break
		 }
		 if target == c {
			res = append(res, []int{target})
			break
		 } else {
		   for _, r := range helper(candidates[i+1:], target - c) {
			 r = append(r, c)
			 res = append(res, r)
		   }
		 }
         i++
	 }
	 return res
}