package stack

type MyStack struct {
	l   int
	arr []int
}

func NewStack() MyStack {
	return MyStack{arr: make([]int, 0, 100)}
}

func (this *MyStack) push(x int) {
	this.arr = append(this.arr, x)
	this.l++
}

func (this *MyStack) pop() int {
	res := this.arr[this.l-1]
	this.arr = this.arr[:this.l-1]
	this.l--
	return res
}

func (this *MyStack) peek() int {
	return this.arr[this.l-1]
}

func (this *MyStack) size() int {
	return this.l
}
func (this *MyStack) isEmpty() bool {
	return this.l == 0
}

type MyQueue struct {
	pushStack MyStack
	popStack  MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{pushStack: NewStack(), popStack: NewStack()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.pushStack.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.popStack.isEmpty() {
		for !this.pushStack.isEmpty() {
			this.popStack.push(this.pushStack.pop())
		}
	}
	return this.popStack.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.popStack.isEmpty() {
		for !this.pushStack.isEmpty() {
			this.popStack.push(this.pushStack.pop())
		}
	}
	return this.popStack.peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.popStack.isEmpty() && this.pushStack.isEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
