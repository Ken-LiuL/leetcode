package array


func countCharacters(words []string, chars string) int {
	chs := make([]int, 26)
	cnt := 0
	for _, c := range chars {
		chs[c - 'a'] += 1
	}
	for _, w := range words {
		rs :=  append([]int(nil), chs...)
		ok := true
		for _, c := range w {
			rs[c - 'a']--
			if  rs[c - 'a'] < 0 {
				ok = false
				break
			}
		}
		if ok {
			cnt += len(w)
		}
	}
}