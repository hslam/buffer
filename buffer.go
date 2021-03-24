// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package buffer

import (
	"sync"
)

const numShards = 256

type buffers struct {
	shards [numShards]Map
}

type Map struct {
	lock sync.RWMutex
	data map[int]*Pool
}

type Pool struct {
	pool *sync.Pool
	size int
}

func (p *Pool) GetBuffer(size int) []byte {
	return p.pool.Get().([]byte)
}

func (p *Pool) PutBuffer(buf []byte) {
	buf = buf[:cap(buf)]
	if cap(buf) >= p.size {
		p.pool.Put(buf)
	}
}

func newBuffers() *buffers {
	b := new(buffers)
	for i := range b.shards {
		b.shards[i].data = make(map[int]*Pool)
	}
	return b
}

func (b *buffers) AssignPool(size int) *Pool {
	m := &b.shards[size%numShards]
	m.lock.RLock()
	if p, ok := m.data[size]; ok {
		m.lock.RUnlock()
		return p
	}
	m.lock.RUnlock()
	m.lock.Lock()
	if p, ok := m.data[size]; ok {
		m.lock.Unlock()
		return p
	}
	p := &Pool{
		pool: &sync.Pool{New: func() interface{} {
			return make([]byte, size)
		}},
		size: size,
	}
	m.data[size] = p
	m.lock.Unlock()
	return p
}

var defaultBuffers = newBuffers()

func AssignPool(size int) *Pool {
	return defaultBuffers.AssignPool(size)
}
