package single_linked_list

import (
	"sync"
)

type Node struct {
	data interface{}
	ptr *Node
}

type SingleLinkedList struct {
	head *Node
	m *sync.RWMutex
}

func (l *SingleLinkedList)Add(inputData interface{}){
	n:=&Node{inputData,nil}
	l.m.Lock()
	defer l.m.Unlock()
	l.head.ptr = n
}