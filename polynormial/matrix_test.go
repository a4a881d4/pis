package polynormal

import (
	"fmt"
	"testing"
)

func TestP128Matrix5X4(t *testing.T) {
	m := make([]int, 20)
	prime := P128[0]
	for tcase := 0; tcase < 256; tcase++ {
		for i := 0; i < 20; i++ {
			m[i] = polyRand.Intn(127)
		}
		g := prime.NewMatrix(5, 4, m)
		A := g.Part(0, 0, 4, 4)
		X := g.Part(4, 0, 1, 4)
		var Y *PMatrix
		var invable bool
		if tcase != 12 {
			Y, invable = g.Guass(false)
		} else {
			Y, invable = g.Guass(false)
		}
		if invable {
			Z := A.Mul(Y)
			if !Z.Equ(X) {
				fmt.Println("Case: ", tcase)
				A.PrintMatrix("A")
				X.PrintMatrix("X")
				Y.PrintMatrix("Y")
				Z.PrintMatrix("Z")
				g.PrintMatrix("g")
				t.Error("P128 Matrix(5X4) Guass, no pass")
			}
		}
	}
}

func TestP128MatrixMul(t *testing.T) {
	m := make([]int, 16)
	for i := 0; i < 16; i++ {
		m[i] = polyRand.Intn(126)
	}
	prime := P128[0]
	A := prime.NewMatrix(4, 4, m)
	y := make([]int64, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			y[i] ^= A.A[i*4+j]
		}
	}
	x := make([]int, 4)
	X := prime.NewMatrix(1, 4, x)
	Y := A.Mul(X)
	if !Y.Equ(prime.NewMatrixInt64(1, 4, y)) {
		t.Error("P128 Mul, no pass")
	}
}
