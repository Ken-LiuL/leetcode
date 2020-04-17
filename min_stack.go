package stack



type Node struct {
	val int
	minval int
    pre  *Node
}
type MinStack struct {
	head  *Node
	l    int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{head: nil, l: 0}
}


func (this *MinStack) Push(x int)  {
	if this.head == nil {
		this.head = &Node{val: x, pre: nil, minval: x}
	} else {
		this.head = &Node{val: x, pre: this.head, minval: min(x, this.head.minval)}
	}
	this.l += 1
}


func (this *MinStack) Pop()  {
   if this.l > 0 {
	  this.head = this.head.pre
	  this.l -= 1
   }
}


func (this *MinStack) Top() int {
	if this.l > 0 {
		return this.head.val
	}
    return 0
}


func (this *MinStack) GetMin() int {
    if this.l > 0 {
		return this.head.minval
	}
    return 0
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */