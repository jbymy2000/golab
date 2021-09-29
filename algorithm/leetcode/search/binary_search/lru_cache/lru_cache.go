package lru_cache

type LRUCache struct {
	size, capacity int
	m              map[int]*LinkNode
	l              *DoubleLinkList
}

type LinkNode struct {
	key, value int
	pre, next  *LinkNode
}

func NewLinkNode(key, value int) *LinkNode {
	return &LinkNode{key: 0, value: 0}
}

type DoubleLinkList struct {
	head, tail *LinkNode
}

func (p *DoubleLinkList) Move2Head(node *LinkNode) {
	node.pre = p.head
	node.next = p.head.next
	p.head.next.pre = node
	p.head.next = node
}

func (p *DoubleLinkList) RemoveTail() *LinkNode {
	currTail := p.tail.pre
	newTail := p.tail.pre.pre
	p.tail.pre = newTail
	newTail.next = p.tail
	return currTail
}

func (p *DoubleLinkList) Remove(node *LinkNode) {
	node.next.pre = node.pre
	node.pre.next = node
}

func NewDoubleLinkList() *DoubleLinkList {
	head := NewLinkNode(0, 0)
	tail := NewLinkNode(0, 0)
	head.next = tail
	tail.pre = head
	return &DoubleLinkList{head: head, tail: tail}
}

func NewLRUCache(capacity int) *LRUCache {
	hashMap := make(map[int]*LinkNode, capacity)
	list := NewDoubleLinkList()
	return &LRUCache{
		size:     0,
		capacity: capacity,
		m:        hashMap,
		l:        list,
	}
}

func (p *LRUCache) Get(key int) (int, bool) {
	var (
		v  *LinkNode
		ok bool
	)
	if v, ok = p.m[key]; ok {
		p.l.Move2Head(v)
		return v.value, true
	}
	return -1, false
}

func (p *LRUCache) Put(key, value int) {
	var (
		v  *LinkNode
		ok bool
	)
	if v, ok = p.m[key]; ok {
		v.value = value
		p.l.Remove(v)
		p.l.Move2Head(v)
		return
	}
	if p.size > p.capacity {
		oldTail := p.l.RemoveTail()
		delete(p.m, oldTail.key)
	}
	newNode := NewLinkNode(key, value)
	p.l.Move2Head(newNode)
}
