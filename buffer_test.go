package buffer

import (
	"testing"
)

func TestAssignPool(t *testing.T) {
	for i := 0; i < 2; i++ {
		size := 1024
		p := AssignPool(size)
		b := p.GetBuffer(size)
		if len(b) < size {
			t.Error(len(b))
		}
	}
}
