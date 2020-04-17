package main

import "fmt"
import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return helper(nums)
}

func helper(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	start := 0
	repeat := 0
	for start < len(nums) - 1 {
		if  nums[start] == nums[start+1] {
			repeat++
			start++
        } else{
            break
        }
	}
	res :=  helper(nums[start+1:])   
    l := len(res)       
	for i:=0; i < l; i++ {
		s := append([]int{}, res[i]...)
		for j := 0; j <= repeat;j++ {
			t := append(s, nums[0])
			res = append(res, append(s, nums[0]))
			s =  append([]int{}, t...)
		}
	}
	return res
}
 