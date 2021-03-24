// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package buffer implements a variable-sized bytes pool.
package buffer

import (
	"sync"
)

const (
	numBuckets = 256
	unit       = 1024
)

type Buffers struct {
	unit    int
	buckets [numBuckets]bucket
}

type bucket struct {
	lock  sync.RWMutex
	pools map[int]*Pool
}

type Pool struct {
	size int
	pool *sync.Pool
}

func (p *Pool) GetBuffer(size int) []byte {
	return p.pool.Get().([]byte)[:size]
}

func (p *Pool) PutBuffer(buf []byte) {
	buf = buf[:cap(buf)]
	if cap(buf) >= p.size {
		p.pool.Put(buf)
	}
}

func NewBuffers(unit int) *Buffers {
	b := new(Buffers)
	for i := range b.buckets {
		b.buckets[i].pools = make(map[int]*Pool)
	}
	b.unit = unit
	return b
}

func (b *Buffers) AssignPool(size int) (p *Pool) {
	if b.unit > 0 && size%b.unit > 0 {
		size = size/b.unit*b.unit + b.unit
	}
	m := &b.buckets[size%numBuckets]
	var ok bool
	m.lock.RLock()
	if p, ok = m.pools[size]; ok {
		m.lock.RUnlock()
		return
	}
	m.lock.RUnlock()
	m.lock.Lock()
	if p, ok = m.pools[size]; !ok {
		p = &Pool{
			pool: &sync.Pool{New: func() interface{} {
				return make([]byte, size)
			}},
			size: size,
		}
		m.pools[size] = p
	}
	m.lock.Unlock()
	return
}

func (b *Buffers) GetBuffer(size int) []byte {
	return b.AssignPool(size).GetBuffer(size)
}

func (b *Buffers) PutBuffer(buf []byte) {
	b.AssignPool(cap(buf)).PutBuffer(buf)
}

var defaultBuffers = NewBuffers(unit)

func AssignPool(size int) *Pool {
	return defaultBuffers.AssignPool(size)
}

func GetBuffer(size int) []byte {
	return defaultBuffers.GetBuffer(size)
}

func PutBuffer(buf []byte) {
	defaultBuffers.PutBuffer(buf)
}
