// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package buffer

import (
	"testing"
)

func TestUnit(t *testing.T) {
	b := NewBuffers(0)
	units := []int{0, 1, 512, 1023, 1024, 1024}
	sizes := []int{1024, 1024, 1024, 1024, 1024, 1025}
	results := []int{1024, 1024, 1024, 2046, 1024, 2048}
	if len(units) != len(sizes) && len(sizes) != len(results) {
		t.Error()
	}
	for i := 0; i < len(units); i++ {
		b.unit = units[i]
		size := sizes[i]
		if b.unit > 0 && size%b.unit > 0 {
			size = size/b.unit*b.unit + b.unit
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

func BenchmarkAssignPoolUnit2(b *testing.B) {
	bs := NewBuffers(2)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit4(b *testing.B) {
	bs := NewBuffers(4)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit8(b *testing.B) {
	bs := NewBuffers(8)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit16(b *testing.B) {
	bs := NewBuffers(16)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit32(b *testing.B) {
	bs := NewBuffers(32)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit64(b *testing.B) {
	bs := NewBuffers(64)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit128(b *testing.B) {
	bs := NewBuffers(128)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit256(b *testing.B) {
	bs := NewBuffers(256)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit512(b *testing.B) {
	bs := NewBuffers(512)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit1024(b *testing.B) {
	bs := NewBuffers(1024)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit2048(b *testing.B) {
	bs := NewBuffers(2048)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit4096(b *testing.B) {
	bs := NewBuffers(4096)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit8192(b *testing.B) {
	bs := NewBuffers(8192)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		bs.AssignPool(size)
	}
}

func BenchmarkAssignPoolUnit16384(b *testing.B) {
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

func BenchmarkBuffersUnit2(b *testing.B) {
	bs := NewBuffers(2)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit4(b *testing.B) {
	bs := NewBuffers(4)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit8(b *testing.B) {
	bs := NewBuffers(8)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit16(b *testing.B) {
	bs := NewBuffers(16)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit32(b *testing.B) {
	bs := NewBuffers(32)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit64(b *testing.B) {
	bs := NewBuffers(64)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit128(b *testing.B) {
	bs := NewBuffers(128)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit256(b *testing.B) {
	bs := NewBuffers(256)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit512(b *testing.B) {
	bs := NewBuffers(512)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit1024(b *testing.B) {
	bs := NewBuffers(1024)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit2048(b *testing.B) {
	bs := NewBuffers(2048)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit4096(b *testing.B) {
	bs := NewBuffers(4096)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit8192(b *testing.B) {
	bs := NewBuffers(8192)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		buf := bs.AssignPool(size).GetBuffer(size)
		bs.AssignPool(size).PutBuffer(buf)
	}
}

func BenchmarkBuffersUnit16384(b *testing.B) {
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

func BenchmarkAssignPoolAndBuffersUnit2(b *testing.B) {
	bs := NewBuffers(2)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit4(b *testing.B) {
	bs := NewBuffers(4)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit8(b *testing.B) {
	bs := NewBuffers(8)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit16(b *testing.B) {
	bs := NewBuffers(16)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit32(b *testing.B) {
	bs := NewBuffers(32)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit64(b *testing.B) {
	bs := NewBuffers(64)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit128(b *testing.B) {
	bs := NewBuffers(128)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit256(b *testing.B) {
	bs := NewBuffers(256)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit512(b *testing.B) {
	bs := NewBuffers(512)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit1024(b *testing.B) {
	bs := NewBuffers(1024)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit2048(b *testing.B) {
	bs := NewBuffers(2048)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit4096(b *testing.B) {
	bs := NewBuffers(4096)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit8192(b *testing.B) {
	bs := NewBuffers(8192)
	for i := 0; i < b.N; i++ {
		size := i % (64 * 1024)
		p := bs.AssignPool(size)
		buf := p.GetBuffer(size)
		p.PutBuffer(buf)
	}
}

func BenchmarkAssignPoolAndBuffersUnit16384(b *testing.B) {
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
