package polynormail

import (
	"testing"
	"math/big"
)

func TestNewXn( t *testing.T ) {
	x := NewXn(1)
	y := big.Int(x)
	(&x).Println("Println")
	(&x).PrintPoly()
	if big.NewInt(2).Cmp(&y) != 0 {
		t.Error("xn = 1, no pass")
	}
}

func TestCopy( t *testing.T ) {
	x := NewRand(256)
	y := big.Int(x)
	(&x).PrintPoly()
	z := NewPoly(&x)
	(&z).PrintPoly()
	iZ := big.Int(z)
	if (&y).Cmp(&iZ) != 0 {
		t.Error("Copy, no pass")
	}
}

func TestAdd( t *testing.T ) {
	x := NewRand(256)
	y := NewRand(256)
	iX, iY := big.Int(x), big.Int(y)
	z := big.NewInt(0)
	z.Xor(&iX,&iY)
	(&x).PrintPoly()
	(&y).PrintPoly()
	pz := NewPoly(&x)
	pz = (&pz).Add(&y)
	iPz := big.Int(pz)
	(&pz).PrintPoly()
	if z.Cmp(&iPz) != 0 {
		t.Error("Add, no pass")
	}
}