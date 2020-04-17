package hash


func isAnagram(s string, t string) bool {
	dict := make([]int, 26)
	for i := 0; i < len(s); i++ {
		dict[s[i] - 'a'] += 1
	}

	for i := 0; i < len(t); i++ {
		dict[t[i] - 'a'] -= 1
	}

	for i := 0; i < 26; i++ {
		if dict[i] != 0 {
			return false
		}
	}
	return true
}