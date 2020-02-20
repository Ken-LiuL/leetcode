package array

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	elems := make([]int, len(arr1))
	for _, v := range arr1 {
		elems[v]++
	}
	var ret []int
	for _, v := range arr2 {
		for elems[v] > 0 {
			ret = append(ret, v)
			elems[v]--
		}
	}
	var a1 []int
	for k, v := range elems {
		for elems[v] > 0 {
			a1 = append(a1, k)
			elems[v]--
		}
	}
	sort.Ints(a1)
	return append(ret, a1...)
}