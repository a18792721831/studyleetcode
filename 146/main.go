package main

// 双向链表
type LRUCache struct {
	size int
	capa int
	head *Node
	tail *Node
	m    map[int]*Node
}

type Node struct {
	k    int
	v    int
	pre  *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capa: capacity,
		size: 0,
		m:    make(map[int]*Node),
	}
	// 双向链表维护
	l.head = &Node{}
	l.tail = &Node{}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if this.size == 0 {
		return -1
	}
	va, ok := this.m[key]
	if !ok {
		return -1
	}
	this.moveToHead(va)
	return va.v
}

func (this *LRUCache) Put(key int, value int) {
	// 先看看有没有，如果有，更新值
	va, ok := this.m[key]
	if ok {
		// 更新值
		va.v = value
		// 更新完值之后，需要移动到最前
		this.moveToHead(va)
	} else {
		// 没有，需要插入
		// 插入总是从头插入
		newNode := &Node{
			k: key,
			v: value,
		}
		this.addHead(newNode)
		this.size++
		// 判断容量是否超出
		if this.size > this.capa {
			// 需要删除尾部
			this.removeTail()
		}
	}
}

func (this *LRUCache) moveToHead(n *Node) {
	this.removeNode(n)
	this.addHead(n)
}

func (this *LRUCache) addHead(n *Node) {
	// 放到 head 后面
	n.next = this.head.next
	n.pre = this.head
	this.head.next.pre = n
	this.head.next = n
	this.m[n.k] = n
}

func (this *LRUCache) removeNode(n *Node) {
	// 找到了，不仅要返回值，还需要处理位置
	// 把 node 移动到最前，删除的时候从最后删
	// 首先把 node 从双向链表里面摘出来
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (this *LRUCache) removeTail() {
	// 处理map
	t := this.tail.pre
	// 需要删除尾部
	this.tail.pre = t.pre
	this.tail.pre.next = this.tail
	delete(this.m, t.k)
	// 处理size
	this.size--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
