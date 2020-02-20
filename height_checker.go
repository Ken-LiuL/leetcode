package  array

func heightChecker(heights []int) int {
	heightToFreq := make([]int, 101)
	
	for _, h := range heights {
		heightToFreq[h]++
	}

	result := 0
	curHeight := 0


	for i:=0; i < len(heights); i++ {
		for heightToFreq[curHeight] == 0 {
			curHeight++
		}
		if curHeight != heights[i] {
			result++
		}
		heightToFreq[curHeight]--
	}
	return result
}

