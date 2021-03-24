// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package buffer

import (
	"testing"
)

func TestAssignPool(t *testing.T) {
	for i := 0; i < 2; i++ {
		size := 1024
		p := AssignPool(size)
		buf := p.GetBuffer(size)
		if len(buf) < size {
			t.Error(len(buf))
		}
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPool(b *testing.B) {
	bs := newBuffers()
	for i := 0; i < b.N; i++ {
		size := i % numShards * 1024
		bs.AssignPool(size)
	}
}

func BenchmarkAssignSizedPool(b *testing.B) {
	bs := newBuffers()
	size := 64 * 1024
	bs.AssignPool(size)
	for i := 0; i < b.N; i++ {
		bs.AssignPool(size)
	}
}

func BenchmarkBuffers(b *testing.B) {
	bs := newBuffers()
	for i := 0; i < b.N; i++ {
		size := i % numShards * 1024
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffers(b *testing.B) {
	bs := newBuffers()
	for i := 0; i < b.N; i++ {
		size := i % numShards * 1024
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkSizedBuffer(b *testing.B) {
	bs := newBuffers()
	length := 64 * 1024
	p := bs.AssignPool(length)
	for i := 0; i < b.N; i++ {
		buf := p.GetBuffer(length)
		p.PutBuffer(buf)
	}
}
