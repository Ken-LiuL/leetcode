package hash

func distributeCandies(candies []int) int {
	cache := make([]bool, 200001)
	kind := 0
	for _, candy := range candies {
		tmp := candy + 100000
		if !cache[tmp] {
			kind++
			cache[tmp] = true
		}
	}
	total := len(candies)
	if total/2 > kind {
		return kind
	} else {
		return total / 2
	}
}
