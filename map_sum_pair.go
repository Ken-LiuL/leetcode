
//another tries data structure
type MapSum struct {
	Children [26]*MapSum
	Val      int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func (root *MapSum) Insert(key string, val int) {
	start := root
	runes := []rune(key)
	for i, r := range runes {
		if node := start.Children[r-'a']; node != nil {
			start = node
		} else {
			for _, nr := range runes[i:] {
				start.Children[nr-'a'] = &MapSum{}
				start = start.Children[nr-'a']
			}
			break
		}
	}
	start.Val = val
}

func (root *MapSum) SumOfChildren() int {
	sum := 0
	for _, c := range root.Children {
		if c != nil {
			sum += c.Val + c.SumOfChildren()
		}
	}
	return sum
}

//downwards, and get sum of all leafs
func (root *MapSum) Sum(prefix string) int {
	start := root
	for _, r := range []rune(prefix) {
		start = start.Children[r-'a']
		if start == nil {
			return 0
		}
	}
	//find the sum of all child nodes
	return start.Val + start.SumOfChildren()
}
