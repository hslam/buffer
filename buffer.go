package buffer

import (
	"sync"
	"sync/atomic"
)

var (
	buffers = sync.Map{}
	assign  int32
)

type Pool struct {
	pool *sync.Pool
	size int
}

func (p *Pool) GetBuffer(length int) []byte {
	return p.pool.Get().([]byte)
}

func (p *Pool) PutBuffer(buf []byte) {
	buf = buf[:cap(buf)]
	if cap(buf) >= p.size {
		p.pool.Put(buf)
	}
}

func AssignPool(size int) *Pool {
	for {
		if p, ok := buffers.Load(size); ok {
			return p.(*Pool)
		}
		if atomic.CompareAndSwapInt32(&assign, 0, 1) {
			var pool = &Pool{
				pool: &sync.Pool{New: func() interface{} {
					return make([]byte, size)
				}},
				size: size,
			}
			buffers.Store(size, pool)
			atomic.StoreInt32(&assign, 0)
			return pool
		}
	}
}
