package hash

import "strings"

func wordPattern(pattern string, str string) bool {
	mapping := make(map[byte]string, 26)
	reverseMapping := make(map[string]byte)
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
        return false
    }
	for i := 0; i < len(pattern); i++ {
		if r, ok := mapping[pattern[i]]; ok {
			if r != strs[i] {
				return false
			}
		} else {
			if _, ok := reverseMapping[strs[i]]; ok {
				return false
			}
			mapping[pattern[i]] = strs[i]
			reverseMapping[strs[i]] = pattern[i]
		}
	}
	return true
}
