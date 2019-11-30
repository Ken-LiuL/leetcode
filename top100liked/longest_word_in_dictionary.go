package main

//still use tries data structure to store all the data
//w is for the string ended in that node
//l is the length of the string
type Node struct {
	Children [26]*Node
	ind      int //store the index of the string, instead of the string
	l        int //store the length of the string
}

func NewNode() *Node {
	return &Node{}
}

func (node *Node) Insert(word string, ind int) {
	start := node
	wl := len(word)
	for i := 0; i < wl; i++ {
		r := word[i]
		if node := start.Children[r-'a']; node != nil {
			start = node
		} else {
			for ; i < wl; i++ {
				next := &Node{}
				start.Children[word[i]-'a'] = next
				start = next
			}
			break
		}
	}
	//word ended in this node
	start.l = wl
	start.ind = ind
}

//the keypoint is to find the longest path, and for each node, the w field is not empty, like:
//                          ()
//                     a         b
//                    b       a  c
//                  c       d  h
//                d           f
//                             d
//
//
//
func (root *Node) Search() *Node {
	if root.l == 0 {
		return root
	}
	cur := root
	for _, node := range root.Children {
		if node != nil {
			res := node.Search()
			if res.l > cur.l {
				cur = res
			}
		}
	}
	return cur
}

func longestWord(words []string) string {
	root := NewNode()
	for i, w := range words {
		root.Insert(w, i)
	}
	root.l = -1
	res := root.Search()
	if res.l > 0 {
		return words[res.ind]
	} else {
		return ""
	}
}
