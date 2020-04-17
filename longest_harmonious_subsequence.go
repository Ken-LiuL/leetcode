package hash


func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func findLHS(nums []int) int {
	cache := make(map[int]int)
	for _, v := range nums {
		cache[v]++
	}
	longest := 0
 	for k, _ := range cache {
		cnt := max(cache[k+1], cache[k-1])
		if cnt == 0 {
			continue
		} 
		cnt += cache[k]
		if cnt > longest {
			longest = cnt
		}
	}
	return longest

}