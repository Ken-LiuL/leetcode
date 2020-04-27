package heap

import "sort"

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	sortedAppend := func(x int, slice []int) []int {
		i := sort.SearchInts(slice,  x)
		return append(slice[:i], append([]int{x}, slice[i:]...)...)
	}

	for  l :=  len(stones); l > 1; l = len(stones) {
		x, y := l-1, l-2
		s := stones[x] -  stones[y]
		stones = stones[:y]
		if s >  0 {
			stones = sortedAppend(s, stones)
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}