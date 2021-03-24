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
