package polynormal

import (
	_ "fmt"
	"testing"
)

func TestP128Matrix5X4(t *testing.T) {
	m := make([]int, 20)
	for i := 0; i < 20; i++ {
		m[i] = polyRand.Intn(126)
	}
	prime := P128[0]
	g := prime.NewMatrix(5, 4, m)
	A := g.Part(0, 0, 4, 4)
	X := g.Part(4, 0, 1, 4)
	A.PrintMatrix("A")
	X.PrintMatrix("X")
	g.Guass()
	Y := g.Part(4, 0, 1, 4)
	Y.PrintMatrix("Y")
	Z := A.Mul(Y)
	Z.PrintMatrix("Z")
	if g.A[0] != 1 {
		t.Error("P128 Matrix(5X4) Guass, no pass")
	}
}
