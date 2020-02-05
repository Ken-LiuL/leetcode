package tree

type Pair struct {
	val *TreeNode
	level int
}
type Queue []*Pair

func newQueue() *Queue {
	return &Queue{}
}
func (q *Queue) enqueue(p *Pair) {
	 *q = append(*q, p)
}
func (q *Queue) dequeue() *Pair {
	tmp :=  (*q)[0]
	*q = (*q)[1:]
	return tmp
}
func (q *Queue) empty() bool {
	return len(*q) == 0
}
//level recursion
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0, 10)
	q := newQueue()
	res = append(res, 0)
	currentLevel := 0
	cnt := 0
	q.enqueue(&Pair{val: root, level: currentLevel})
	for !q.empty() {
		e := q.dequeue()
		r, l := e.val, e.level
		if r.Left != nil {
			q.enqueue(&Pair{val: r.Left, level: l + 1})
		}
		if r.Right != nil {
			q.enqueue(&Pair{val: r.Right, level: l + 1})
		}
		if l == currentLevel {
			res[currentLevel] += float64(r.Val)
			cnt++
		} else {
			res[currentLevel] = res[currentLevel] / float64(cnt)
			res = append(res, r.Val)
			currentLevel = l
			cnt=1
		}
	}
	res[currentLevel] = res[currentLevel] / float64(cnt)
	return res
}