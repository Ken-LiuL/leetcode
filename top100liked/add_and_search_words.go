package main

type WordDictionary struct {
	Children [27]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (root *WordDictionary) AddWord(word string) {
	start := root
	strLen := len(word)
	for i := 0; i < strLen; i++ {
		if node := start.Children[word[i]-'a']; node != nil {
			start = node
		} else {
			for i < strLen {
				next := &WordDictionary{}
				start.Children[word[i]-'a'] = next
				start = next
				i++
			}
			break
		}
	}
	if start.Children[26] == nil {
		start.Children[26] = &WordDictionary{}
	}
}

func (root *WordDictionary) SearchHelper(word []rune, ind int) bool {
	start := root
	strLen := len(word)

	for i := ind; i < strLen; i++ {
		r := word[i]
		if r == '.' {
			for _, c := range start.Children {
				if c != nil {
					if c.SearchHelper(word, i+1) {
						return true
					}
				}
			}
			return false

		} else if node := start.Children[r-'a']; node != nil {
			start = node
		} else {
			return false
		}
	}

	if start.Children[26] != nil {
		return true
	}
	return false
}

func (root *WordDictionary) Search(word string) bool {
	//rune slice is faster than passing string 
	return root.SearchHelper([]rune(word), 0)
}
