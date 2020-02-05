package main

//use tries to store bits of number, O(32N)
//and then search maximum match for each one, O(N)
//totally, it is O(N)
type Digit struct {
	Children [2]*Digit
	Val      int
}

func NewDigit() *Digit {
	return &Digit{}
}

func (root *Digit) Insert(number int) {
	start := root
	for i := 31; i >= 0; i-- {
		b := (number >> uint(i)) & 1
		if start.Children[b] == nil {
			start.Children[b] = &Digit{}
		}
		start = start.Children[b]
	}
	start.Val = number
}

func (root *Digit) FindMaximumMatch(num int) int {
	start := root
	for i := 31; i >= 0; i-- {
		b := (num >> uint(i)) & 1
		if node := start.Children[b^1]; node != nil {
			start = node
		} else if node := start.Children[b]; node != nil {
			start = node
		} else {
			return start.Val
		}
	}
	return start.Val
}

func findMaximumXOR(nums []int) int {
	tries := NewDigit()
	for _, n := range nums {
		tries.Insert(n)
	}
	largest := -1
	for _, n := range nums {
		m := tries.FindMaximumMatch(n)
		cur := n ^ m
		if cur > largest {
			largest = cur
		}
	}
	return largest
}
