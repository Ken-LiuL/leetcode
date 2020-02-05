package main

import "strings"

type DictCharacter struct {
	Children [26]*DictCharacter
	W        string
}

func NewDictCharacter() *DictCharacter {
	return &DictCharacter{}
}

func (root *DictCharacter) Insert(word string) {
	start := root
	for i, w := range word {
		if node := start.Children[w-'a']; node != nil {
			start = node
		} else {
			for _, nw := range word[i:] {
				n := &DictCharacter{}
				start.Children[nw-'a'] = n
				start = n
			}
			break
		}
	}
	start.W = word
}

func (root *DictCharacter) SearchShortestPath(word string, ind int, wl int) string {

	if wl == ind {
		return root.W
	}
	if root.W != "" {
		return root.W
	}
	c := root.Children[word[ind]-'a']
	if c != nil {
		return c.SearchShortestPath(word, ind+1, wl)
	} else {
		return root.W
	}
}

func replaceWords(dict []string, sentence string) string {
	d := NewDictCharacter()
	words := strings.Split(sentence, " ")
	for _, w := range dict {
		d.Insert(w)
	}
	for i, s := range words {
		r := d.SearchShortestPath(s, 0, len(s))
		if r != "" {
			words[i] = r
		}
	}
	return strings.Join(words, " ")
}
