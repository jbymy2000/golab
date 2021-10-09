package v2

type LinkNode struct {
	key, value int
	pre, next  *LinkNode
}

type LinkList struct {
	head, tail *LinkNode
}

func NewLinkList() *LinkList {
	head := &LinkNode{key: 0, value: 0}
	tail := &LinkNode{key: 0, value: 0}
	head.next = tail
	tail.pre = head
	return &LinkList{head: head, tail: tail}
}

func (p *LinkList) MoveNode2Head(node *LinkNode) {
	p.RemoveNode(node)
	p.AddNode2Head(node)
}

func (p *LinkList) RemoveTail() int {
	key := p.tail.key
	p.tail.next.pre = p.tail.pre
	p.tail.pre.next = p.tail.next
	p.tail = p.tail.pre
	return key
}

func (p *LinkList) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (p *LinkList) AddNode2Head(node *LinkNode) {
	node.pre = p.head
	node.next = p.head.next
	node.next.pre = node
	p.head.next = node
}

type LRUCache struct {
	size, capacity int
	l              *LinkList
	m              map[int]*LinkNode
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		size:     0,
		m:        make(map[int]*LinkNode),
		l:        NewLinkList(),
	}
}

func (p *LRUCache) Get(key int) (value int, ok bool) {
	var (
		node *LinkNode
	)
	if node, ok = p.m[key]; ok {
		p.l.MoveNode2Head(node)
		return node.value, true
	}
	return -1, false
}

func (p *LRUCache) Set(key, value int) {
	var (
		queryNode *LinkNode
		ok        bool
	)

	if queryNode, ok = p.m[key]; ok {
		queryNode.value = value
		p.l.MoveNode2Head(queryNode)
		return
	}

	if p.size > p.capacity {
		tailKey := p.l.RemoveTail()
		delete(p.m, tailKey)
		p.size--
	}
	newNode := &LinkNode{key: key, value: value}
	p.l.AddNode2Head(newNode)
	p.m[key] = newNode
	p.size++
}
