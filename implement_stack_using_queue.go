package stack

type MyQueue struct {
	size int
	arr  []int
}

func NewQueue() MyQueue {
	return MyQueue{arr: make([]int, 0, 100)}
}

func (this *MyQueue) push(x int) {
	this.arr = append(this.arr, x)
	this.size++
}
func (this *MyQueue) pop() int {
	res := this.arr[0]
	this.arr = this.arr[1:]
	this.size--
	return res
}
func (this *MyQueue) peek() int {
	return this.arr[0]
}
func (this *MyQueue) isEmpty() bool {
	return this.size == 0
}
func (this *MyQueue) Qsize() int {
	return this.size
}

type MyStack struct {
	current MyQueue
	backup  MyQueue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{current: NewQueue(), backup: NewQueue()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.current.push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.current.Qsize() != 1 {
		this.backup.push(this.current.pop())
	}
	res := this.current.pop()
	this.current, this.backup = this.backup, this.current

	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for this.current.Qsize() != 1 {
		this.backup.push(this.current.pop())
	}
	res := this.current.pop()
	this.backup.push(res)
	this.current, this.backup = this.backup, this.current

	return res
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.current.isEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
