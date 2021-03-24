// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package buffer

import (
	"sync"
	"sync/atomic"
)

type buffers struct {
	sync.Map
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
	return new(buffers)
}

func (b *buffers) AssignPool(size int) *Pool {
	for {
		if p, ok := b.Load(size); ok {
			return p.(*Pool)
		}
		if atomic.CompareAndSwapInt32(&b.assign, 0, 1) {
			var pool = &Pool{
				pool: &sync.Pool{New: func() interface{} {
					return make([]byte, size)
				}},
				size: size,
			}
			b.Store(size, pool)
			atomic.StoreInt32(&b.assign, 0)
			return pool
		}
	}
}

var defaultBuffers = newBuffers()

func AssignPool(size int) *Pool {
	return defaultBuffers.AssignPool(size)
}
