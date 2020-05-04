package string

func canConstruct(ransomNote string, magazine string) bool {
	dict := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		dict[magazine[i] - 'a']++
	}

	for i := 0; i < len(ransomNote);i++ {
		dict[ransomNote[i] - 'a']--
		if dict[ransomNote[i] - 'a'] < 0 {
			return false
		}
	}
	return true

}