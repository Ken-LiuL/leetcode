 
//FIFO queue
type Queue []*Pair
func newQueue() *Queue {
	return &Queue{}
}

func (q *Queue) inqueue(n *Pair) {
	*q = append(*q, n)
}

func (q *Queue) outqueue() *Pair {
    var tmp *Pair
	if !q.isEmpty() {
        tmp = (*q)[0]
        *q = (*q)[1:]
	}
    return tmp
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

//store TreeNode an relevant level value
type Pair struct {
	_1 *TreeNode
	_2 int
}

func levelOrder(root *TreeNode) [][]int {
	 res := make([][]int, 0, 20)
	 if root == nil {
		 return res
	 }
	 queue := newQueue()
	 //store node and level
	 queue.inqueue(&Pair{root, 0})
	 for !queue.isEmpty() {
		 ele := queue.outqueue()
		 n, level := ele._1, ele._2
		 //if no such for this level
		 if len(res) < level+1 {
			 res = append(res, make([]int, 0, 10))
		 }
		 //append value 
		 res[level] = append(res[level], n.Val)
		 if n.Left != nil {
			queue.inqueue(&Pair{n.Left, level+1})
		 }
		 if n.Right != nil {
			queue.inqueue(&Pair{n.Right, level+1})
		 }
	 }
	return res
}