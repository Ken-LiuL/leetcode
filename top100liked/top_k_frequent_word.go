package main

type CharacterNode struct {
	Children [26]*CharacterNode
	Freq     int
	W        string
}

func NewCharacterNode() *CharacterNode {
	return &CharacterNode{}
}
func (root *CharacterNode) Insert(word string) {
	start := root
	runes := []rune(word)
	for i, r := range runes {
		if node := start.Children[r-'a']; node != nil {
			start = node
		} else {
			for _, nr := range runes[i:] {
				start.Children[nr-'a'] = &CharacterNode{}
				start = start.Children[nr-'a']
			}
			break
		}
	}
	start.Freq++
	start.W = word
}

func (root *CharacterNode) Tranverse(buckets [][]string) {
	for _, c := range root.Children {
		if c != nil {
			if c.Freq != 0 {
				if buckets[c.Freq] == nil {
					buckets[c.Freq] = make([]string, 0, wordsLen)
				}
				buckets[c.Freq] = append(buckets[c.Freq], c.W)
			}
			c.Tranverse(buckets)
		}
	}
}

var wordsLen int

//tries and bucket sort
//tries is to make sure string are retrieved in alphabetical order
func topKFrequent(words []string, k int) []string {
	wordsLen = len(words)
	bucket := make([][]string, wordsLen, wordsLen)
	res := make([]string, 0, k)

	dict := NewCharacterNode()
	for _, w := range words {
		dict.Insert(w)
	}
	s := k
	//put words alphabetically in buckets
	dict.Tranverse(bucket)
	//get tgo rop k elements from the buckets
	for i := wordsLen - 1; i >= 0 && k > 0; i-- {
		if bucket[i] != nil {
			k -= len(bucket[i])
			for _, b := range bucket[i] {
				res = append(res, b)
			}
		}
	}
	return res[:s]
}
