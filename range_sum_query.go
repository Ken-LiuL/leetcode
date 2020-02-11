package dynamic

type NumArray struct {
	cache map[int]int
}

func Constructor(nums []int) NumArray {
	na := NumArray{cache: make(map[int]int)}
	res := 0
	na.cache[0] = 0
	for ind, n  := range nums {
		res += n
		na.cache[ind+1] = res
	} 
	return na
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.cache[j+1] - this.cache[i]
}