package gopool

import (
	"sync"
)

const _default_pool_cap = 1024
type GoPool struct {
	sync.Mutex
	pool []interface{}
	size int
	capacity int
	creater func() interface{}
}
//type func ObjCreater() interface{}
func NewGoPool(capacity int, creater func() interface{}) *GoPool {
	if capacity <= 0 {
		capacity = _default_pool_cap
	}	
	ret := new(GoPool)
	ret.pool = make([]interface{}, capacity)
	ret.size = 0
	ret.capacity = capacity
	ret.creater = creater
	return ret
}

func (m *GoPool) Get() interface{} {
	m.Lock()
	defer m.Unlock()
	if m.size > 0 {
		m.size--
		ret := m.pool[m.size]
		m.pool[m.size] = nil
		return ret
	} else {
		if m.creater != nil {
			return m.creater()
		} else {
			return nil
		}
	}
}

func (m *GoPool) Put(x interface{}) {
	m.Lock()
	defer m.Unlock()
	if x ==  nil || m.size >= m.capacity {
		return
	}
	m.pool[m.size] = x
	m.size++
	return
}

func (m *GoPool) Size() int {
	return m.size
}



