// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package buffer

import (
	"testing"
)

func TestAssignPool(t *testing.T) {
	buffers := NewBuffers(0)
	if buffers.pageSize != minPageSize {
		t.Errorf("should be %d, but got %d", minPageSize, buffers.pageSize)
	}
	defaultBuffers = NewBuffers(1024)
	for i := 0; i < 65*1024; i++ {
		size := i
		p := AssignPool(size)
		if p.size < size {
			t.Error(p.size)
		}
		buf := GetBuffer(size)
		if len(buf) < size {
			t.Error(len(buf))
		}
		PutBuffer(buf)
	}
}

func BenchmarkAssignPool(b *testing.B) {
	bs := NewBuffers(1024)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignSizedPool(b *testing.B) {
	bs := NewBuffers(1024)
	size := 64 * 1024
	bs.AssignPool(size)
	for i := 0; i < b.N; i++ {
		bs.AssignPool(size)
	}
}

func BenchmarkBuffers(b *testing.B) {
	bs := NewBuffers(1024)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkSizedBuffer(b *testing.B) {
	bs := NewBuffers(1024)
	size := 64 * 1024
	p := bs.AssignPool(size)
	for i := 0; i < b.N; i++ {
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}
