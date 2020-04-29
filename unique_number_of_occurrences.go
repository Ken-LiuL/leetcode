package hash

func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	unique_occur := make(map[int]int)

	for _, n := range arr {
		occurrences[n]++
	}

	for _, occu := range occurrences {
		if _, ok := unique_occur[occu]; ok {
			return false
		} else {
			unique_occur[occu] = 1
		}
	}
	return true
}
