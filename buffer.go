// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package buffer

import (
	"sync"
	"sync/atomic"
)

const numShards = 256

type buffers struct {
	shards [numShards]Map
}

type Map struct {
	*sync.Map
	assign int32
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
		b.shards[i].Map = &sync.Map{}
	}
	return b
}

func (b *buffers) AssignPool(size int) *Pool {
	m := &b.shards[size%numShards]
	for {
		if p, ok := m.Load(size); ok {
			return p.(*Pool)
		}
		if atomic.CompareAndSwapInt32(&m.assign, 0, 1) {
			var pool = &Pool{
				pool: &sync.Pool{New: func() interface{} {
					return make([]byte, size)
				}},
				size: size,
			}
			m.Store(size, pool)
			atomic.StoreInt32(&m.assign, 0)
			return pool
		}
	}
}

var defaultBuffers = newBuffers()

func AssignPool(size int) *Pool {
	return defaultBuffers.AssignPool(size)
}
