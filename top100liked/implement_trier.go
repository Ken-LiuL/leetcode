package main

//Trie represent a node and its children: <character, Trie>
type Trie struct {
	//at first , []*Trie slice is used, but when I change from slice to array, the speed doubled,
	//reason is that, to instantiate a slice, it will also instantiate an array underneath
	Children [27]*Trie
}

//Constructor return a empty root node
func Constructor() Trie {
	return Trie{}
}

//Insert put the word into the tier
func (root *Trie) Insert(word string) {
	start := root
	for i, r := range word {
		if node := start.Children[r-'a']; node != nil {
			start = node
		} else {
			for _, nr := range word[i:] {
				next := &Trie{}
				start.Children[nr-'a'] = next
				start = next
			}
			break
		}
	}
	//add '$' to the end of string , so that all words are leaf
	start.Children[26] = &Trie{}
}

//Search search the word from root to leaf, when verify whether it stops at leaf
func (root *Trie) Search(word string) bool {
	start := root
	for _, r := range word {
		if node := start.Children[r-'a']; node != nil {
			start = node
		} else {
			return false
		}
	}
	if node := start.Children[26]; node != nil {
		return true
	}
	return false
}

//StartsWith the logic is the same with Search, except no need to check whether ends up leaf
func (root *Trie) StartsWith(prefix string) bool {
	start := root
	for _, r := range prefix {
		if node := start.Children[r-'a']; node != nil {
			start = node
		} else {
			return false
		}
	}
	return true
}
