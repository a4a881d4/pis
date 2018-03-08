package polynormal

import (
	_ "fmt"
	"testing"
)

func TestP128Solution5X4(t *testing.T) {
	m := make([]int, 20)
	for i := 0; i < 20; i++ {
		m[i] = polyRand.Intn(127)
	}
	base := NewPolyBase(P128)
	b, x := base.Solution(5, 4, m)
}
