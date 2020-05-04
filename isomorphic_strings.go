package hash

func isIsomorphic(s string, t string) bool {
	mapping := make(map[byte]byte, 256)
	reverseMapping := make(map[byte]byte, 256)
	
	for i := 0; i < len(s); i++ {
		c1, c2 := s[i]-'a', t[i]-'a'
		if r, ok := mapping[c1]; ok {
			if c2 != r {
				return false
			}
		} else {
			if _, ok := reverseMapping[c2]; ok {
				return false
			}
			mapping[c1] = c2
			reverseMapping[c2] = c1
		}
	}
	return true
}
