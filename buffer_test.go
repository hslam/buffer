package buffer

import (
	"testing"
)

func TestAssignPool(t *testing.T) {
	for i := 0; i < 2; i++ {
		p := AssignPool(1024)
		b := p.Get().([]byte)
		if len(b) < 1024 {
			t.Error(len(b))
		}
	}
}
