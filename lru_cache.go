package list

type CacheNode struct {
		prev *CacheNode
		next *CacheNode
		key int
		val int
}
type LRUCache struct {
		cap int
		len int
		head *CacheNode
		tail *CacheNode
		mapping map[int]*CacheNode
}

func Constructor(capacity int) LRUCache {
	 head := &CacheNode{}
	 tail := &CacheNode{} 
	 tail.prev, tail.next, head.prev, head.next = head, head, tail, tail
	 return LRUCache{cap: capacity, 
		mapping: make(map[int]*CacheNode, capacity),
		head: head,
		tail: tail
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.mapping[key]; ok {
		 val := node.val
		 //move to thead
		 this.moveToHead(node)
		 return val
	} 
	return  -1
}

func (this *LRUCache) Put(key int, value int)  {
	found, ok := this.mapping[key]
	if ok {
		found.val = value
		this.moveToHead(found)
		return 
	}
	this.mapping[key] = &CacheNode{key: key, val: value}
	//else check capacity
	if this.cap > this.len {
		this.len++
	} else {
		//delete least usage
		this.removeLeast()
	}
	this.insertToHead(this.mapping[key])
}

func (this *LRUCache) moveToHead(node *CacheNode) {
   node.prev.next, node.next.prev = node.next,  node.prev

   this.head.next.prev, node.next = node, this.head.next
   this.head.next, node.prev = node, this.head
}
func (this *LRUCache) insertToHead(node *CacheNode) {
	this.head.next.prev, node.next = node, this.head.next
	this.head.next, node.prev = node, this.head
}
func (this *LRUCache) removeLeast() {
	//remove  tail 
	least := this.tail.prev
	this.tail.prev.prev.next, this.tail.prev = this.tail, this.tail.prev.prev
	delete(this.mapping, least.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */