// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package buffer implements a variable-sized bytes pool.
package buffer

import (
	"sync"
	"unsafe"
)

const (
	minPageSize = int(unsafe.Sizeof(int(0)))
	threshold   = 64 * 1024
	numBuckets  = 256
	pageSize    = 1024
)

// Buffers contains buckets for sharding.
type Buffers struct {
	pageSize int
	pools    [61]*Pool
	buckets  [numBuckets]bucket
}

type bucket struct {
	lock  sync.RWMutex
	pools map[int]*Pool
}

// Pool represents a fixed size bytes pool.
type Pool struct {
	size int
	pool *sync.Pool
}

// GetBuffer returns a bytes from the pool with the given size.
func (p *Pool) GetBuffer(size int) (buf []byte) {
	buf = p.pool.Get().([]byte)
	if size > 0 {
		buf = buf[:size]
	}
	return
}

// PutBuffer frees the bytes to the pool.
func (p *Pool) PutBuffer(buf []byte) {
	buf = buf[:cap(buf)]
	if cap(buf) >= p.size {
		p.pool.Put(buf)
	}
}

// NewBuffers creates a new Buffers with the given page size.
func NewBuffers(pageSize int) *Buffers {
	b := new(Buffers)
	for i := range b.pools {
		alignedSize := uint(1) << uint(i+3)
		b.pools[i] = &Pool{
			size: int(alignedSize),
			pool: &sync.Pool{New: func() interface{} {
				return make([]byte, alignedSize)
			}},
		}
	}
	for i := range b.buckets {
		b.buckets[i].pools = make(map[int]*Pool)
	}
	if pageSize < minPageSize {
		b.pageSize = minPageSize
	} else {
		b.pageSize = pageSize
	}
	return b
}

// AssignPool assigns a fixed size bytes pool with the given size.
func (b *Buffers) AssignPool(size int) *Pool {
	if size < threshold {
		return b.pools[assignIndex(size)]
	}
	var alignedSize = size
	if size%b.pageSize > 0 {
		alignedSize = size/b.pageSize*b.pageSize + b.pageSize
	}
	m := &b.buckets[alignedSize/b.pageSize%numBuckets]
	var ok bool
	m.lock.RLock()
	var p *Pool
	if p, ok = m.pools[alignedSize]; ok {
		m.lock.RUnlock()
		return p
	}
	m.lock.RUnlock()
	m.lock.Lock()
	if p, ok = m.pools[alignedSize]; !ok {
		p = &Pool{
			pool: &sync.Pool{New: func() interface{} {
				return make([]byte, alignedSize)
			}},
			size: alignedSize,
		}
		m.pools[alignedSize] = p
	}
	m.lock.Unlock()
	return p
}

// GetBuffer returns a bytes from the pool with the given size.
func (b *Buffers) GetBuffer(size int) []byte {
	return b.AssignPool(size).GetBuffer(size)
}

// PutBuffer frees the bytes to the pool.
func (b *Buffers) PutBuffer(buf []byte) {
	b.AssignPool(cap(buf)).PutBuffer(buf)
}

func assignIndex(size int) (idx uint) {
	var t = size
	switch {
	case t < 1<<3+1:
		return 0
	case t < 1<<4+1:
		return 1
	case t < 1<<5+1:
		return 2
	case t < 1<<6+1:
		return 3
	case t < 1<<7+1:
		return 4
	case t < 1<<8+1:
		return 5
	case t < 1<<9+1:
		return 6
	case t < 1<<10+1:
		return 7
	case t < 1<<11+1:
		return 8
	case t < 1<<12+1:
		return 9
	case t < 1<<13+1:
		return 10
	case t < 1<<14+1:
		return 11
	case t < 1<<15+1:
		return 12
	}
	return 13
}

// defaultBuffers is the default instance of *Buffers.
var defaultBuffers = NewBuffers(pageSize)

// AssignPool assigns a fixed size bytes pool with the given size.
func AssignPool(size int) *Pool {
	return defaultBuffers.AssignPool(size)
}

// GetBuffer returns a bytes from the pool with the given size.
func GetBuffer(size int) []byte {
	return defaultBuffers.GetBuffer(size)
}

// PutBuffer frees the bytes to the pool.
func PutBuffer(buf []byte) {
	defaultBuffers.PutBuffer(buf)
}
