package hash

func numJewelsInStones(J string, S string) int {
	dict := make(map[rune]int)
	cnt := 0
	for _, r := range J {
		dict[r] = 1
	}

	for _, r := range S {
		if _, ok := dict[r]; ok {
			cnt++
		}
	}

	return cnt

}
