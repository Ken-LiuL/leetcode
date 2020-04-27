package array

const MaxInt = int(^uint(0) >> 1)
func minimumAbsDifference(arr []int) [][]int {
	res := make([][]int, 0, 10)
	sort.Ints(arr)

	m :=  MaxInt
	for i := 0; i < len(arr)-1; i++ {
		v  := arr[i+1] - arr[i]
		if v < m {
			m = v
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] - arr[i] == m {
			res =  append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}