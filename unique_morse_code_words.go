package string


var mapping []string = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

func uniqueMorseRepresentations(words []string) int {
	cache := make(map[string]bool)

	for  _, s := range words {
		 morse := ""
		 for i := 0; i < len(s); i++ {
			 morse = morse + mapping[s[i] - 'a']
		 }
		 cache[morse] = true
	}
	return len(cache)
}