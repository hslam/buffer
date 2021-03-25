// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package buffer

import (
	"testing"
)

func TestPage(t *testing.T) {
	b := NewBuffers(0)
	pages := []int{0, 1, 512, 1023, 1024, 1024}
	sizes := []int{1024, 1024, 1024, 1024, 1024, 1025}
	results := []int{1024, 1024, 1024, 2046, 1024, 2048}
	if len(pages) != len(sizes) && len(sizes) != len(results) {
		t.Error()
	}
	for i := 0; i < len(pages); i++ {
		b.page = pages[i]
		size := sizes[i]
		if b.page > 0 && size%b.page > 0 {
			size = size/b.page*b.page + b.page
		}
		if size != results[i] {
			t.Error(i, size, results[i])
		}
	}
}

func TestAssignPool(t *testing.T) {
	defaultBuffers = NewBuffers(1024)
	for i := 0; i < 4; i++ {
		size := 64*1024 + i
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
	bs := NewBuffers(0)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage2(b *testing.B) {
	bs := NewBuffers(2)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage4(b *testing.B) {
	bs := NewBuffers(4)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage8(b *testing.B) {
	bs := NewBuffers(8)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage16(b *testing.B) {
	bs := NewBuffers(16)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage32(b *testing.B) {
	bs := NewBuffers(32)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage64(b *testing.B) {
	bs := NewBuffers(64)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage128(b *testing.B) {
	bs := NewBuffers(128)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage256(b *testing.B) {
	bs := NewBuffers(256)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage512(b *testing.B) {
	bs := NewBuffers(512)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage1024(b *testing.B) {
	bs := NewBuffers(1024)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage2048(b *testing.B) {
	bs := NewBuffers(2048)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage4096(b *testing.B) {
	bs := NewBuffers(4096)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage8192(b *testing.B) {
	bs := NewBuffers(8192)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolPage16384(b *testing.B) {
	bs := NewBuffers(16384)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignSizedPool(b *testing.B) {
	bs := NewBuffers(0)
	size := 64 * 1024
	bs.AssignPool(size)
	for i := 0; i < b.N; i++ {
		bs.AssignPool(size)
	}
}

func BenchmarkBuffers(b *testing.B) {
	bs := NewBuffers(0)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage2(b *testing.B) {
	bs := NewBuffers(2)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage4(b *testing.B) {
	bs := NewBuffers(4)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage8(b *testing.B) {
	bs := NewBuffers(8)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage16(b *testing.B) {
	bs := NewBuffers(16)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage32(b *testing.B) {
	bs := NewBuffers(32)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage64(b *testing.B) {
	bs := NewBuffers(64)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage128(b *testing.B) {
	bs := NewBuffers(128)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage256(b *testing.B) {
	bs := NewBuffers(256)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage512(b *testing.B) {
	bs := NewBuffers(512)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage1024(b *testing.B) {
	bs := NewBuffers(1024)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage2048(b *testing.B) {
	bs := NewBuffers(2048)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage4096(b *testing.B) {
	bs := NewBuffers(4096)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage8192(b *testing.B) {
	bs := NewBuffers(8192)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersPage16384(b *testing.B) {
	bs := NewBuffers(16384)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffers(b *testing.B) {
	bs := NewBuffers(0)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage2(b *testing.B) {
	bs := NewBuffers(2)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage4(b *testing.B) {
	bs := NewBuffers(4)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage8(b *testing.B) {
	bs := NewBuffers(8)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage16(b *testing.B) {
	bs := NewBuffers(16)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage32(b *testing.B) {
	bs := NewBuffers(32)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage64(b *testing.B) {
	bs := NewBuffers(64)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage128(b *testing.B) {
	bs := NewBuffers(128)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage256(b *testing.B) {
	bs := NewBuffers(256)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage512(b *testing.B) {
	bs := NewBuffers(512)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage1024(b *testing.B) {
	bs := NewBuffers(1024)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage2048(b *testing.B) {
	bs := NewBuffers(2048)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage4096(b *testing.B) {
	bs := NewBuffers(4096)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage8192(b *testing.B) {
	bs := NewBuffers(8192)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersPage16384(b *testing.B) {
	bs := NewBuffers(16384)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkSizedBuffer(b *testing.B) {
	bs := NewBuffers(0)
	size := 64 * 1024
	p := bs.AssignPool(size)
	for i := 0; i < b.N; i++ {
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}
