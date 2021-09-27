package go_concurrency

import "sync"

type m struct {
	data map[string]string
	lock sync.RWMutex
}

func NewM() *m {
	return &m{data: make(map[string]string)}
}

func (m *m) Set(key, value string) {
	m.lock.Lock()
	m.data[key] = value
	m.lock.Unlock()
}
