package list

type StockSpanner struct {
	q []int   
	v []int
	top int
}


func Constructor() StockSpanner {
    return StockSpanner{q: []int{}, v: []int{}}
}

func (this *StockSpanner) push(v int) {
	 this.q = append(this.q, v)
	 this.top++
}
func (this *StockSpanner) pop() int {
	 res := this.q[this.top - 1]
	 this.q = this.q[0:this.top - 1]
	 this.top--
	 return res
}
func (this *StockSpanner) peek() int {
	if this.top == 0 {
		return -1
	}
	return this.q[this.top - 1]
}

func (this *StockSpanner) Next(price int) int {
	this.v = append(this.v, price)
	for len(this.q) != 0 &&  this.v[this.peek()] <= price {
		this.pop()
	}
	res := len(this.v)-1 - this.peek()
	this.push(len(this.v) - 1)
	return res
}