package main

import "container/list"

// 双向链表
type LRUCache1 struct {
	capa int
	l    *list.List
	m    map[int]*list.Element
}

type ent struct {
	k int
	v int
}

func Constructor1(capacity int) LRUCache1 {
	l := LRUCache1{
		capa: capacity,
		l:    list.New(),
		m:    make(map[int]*list.Element),
	}
	return l
}

func (this *LRUCache1) Get(key int) int {
	va, ok := this.m[key]
	if !ok {
		return -1
	}
	this.l.MoveToFront(va)
	return va.Value.(*ent).v
}

func (this *LRUCache1) Put(key int, value int) {
	// 先看看有没有，如果有，更新值
	va, ok := this.m[key]
	if ok {
		// 更新完值之后，需要移动到最前
		this.l.MoveToFront(va)
		// 更新值
		va.Value = &ent{k: key, v: value}
	} else {
		// 没有，需要插入
		// 插入总是从头插入
		e := &ent{k: key, v: value}
		this.m[key] = this.l.PushFront(e)
		// 判断容量是否超出
		if this.l.Len() > this.capa {
			// 需要删除尾部
			last := this.l.Back()
			this.l.Remove(last)
			delete(this.m, last.Value.(*ent).k)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
