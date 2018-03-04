package polynormail

import (
	"math/big"
	"testing"
)

func TestNewXn(t *testing.T) {
	x := NewXn(1)
	y := x.p
	x.Println("Println")
	x.PrintPoly()
	if big.NewInt(2).Cmp(&y) != 0 {
		t.Error("xn = 1, no pass")
	}
}

func TestCopy(t *testing.T) {
	x := NewRand(256)
	y := x.p
	x.PrintPoly()
	z := NewPoly(x)
	z.PrintPoly()
	iZ := z.p
	if (&y).Cmp(&iZ) != 0 {
		t.Error("Copy, no pass")
	}
}

func TestAdd(t *testing.T) {
	x := NewRand(256)
	y := NewRand(256)
	iX, iY := x.p, y.p
	z := big.NewInt(0)
	z.Xor(&iX, &iY)
	x.PrintPoly()
	y.PrintPoly()
	pz := NewPoly(x)
	az := pz.Add(pz, y)
	iPz := az.p
	pz.PrintPoly()
	if z.Cmp(&iPz) != 0 {
		t.Error("Add, no pass")
	}
}

func TestMul(t *testing.T) {
	x := NewXn(1)
	x.PrintPoly()
	y := NewXn(3)
	y.PrintPoly()
	z := NewXn(4)
	z.PrintPoly()
	x.Mul(x, y)
	x.PrintPoly()
	if (&z.p).Cmp(&x.p) != 0 {
		t.Error("Mul, no pass")
	}
}
