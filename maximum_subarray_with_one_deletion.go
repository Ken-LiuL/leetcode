package dynamic

const MIN_V = -100000
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//the point is:
//maxSkip =  max(last-no-skip, or maxSkip no skip of current value)
func maximumSum(arr []int) int {
	maxSkip, maxNoSkip, maxGlobal := MIN_V, MIN_V, MIN_V
    [1,-2,0,3]
	for _, v := range arr {
       maxSkip = max(maxNoSkip, maxSkip+v)
	   maxNoSkip = max(maxNoSkip+ v, v)
	   maxGlobal = max(max(maxNoSkip, maxSkip), maxGlobal)
	}
	return maxGlobal
}
