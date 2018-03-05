package polynormal

import (
	"fmt"
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
	z := x.NewPoly()
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
	pz := x.NewPoly()
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

func TestDivRem(t *testing.T) {
	x := NewRand(20)
	x.Println("x")
	y := NewRand(20)
	y.Println("y")
	z := NewRand(5)
	z.Println("z")
	s := NewXn(0)
	s.Add(s.Mul(x, y), z)
	s.Println("s")
	f := s.DivRem(y)
	if (&z.p).Cmp(&s.p) != 0 {
		t.Error("Rem, no pass")
	}
	if (&f.p).Cmp(&x.p) != 0 {
		f.Println("d")
		x.Println("x")
		t.Error("Div, no pass")
	}
}

func TestFindPrime(t *testing.T) {
	r := FindPrime(10)
	x := NewXn(0)
	for _, d := range r {
		o := d.Order()
		if o == 2 || o == 8 {
			x.Mul(x, d)
		}
	}
	x.PrintPoly()
	xt := NewXn(0)
	xt.Add(NewXn(128), NewXn(1))
	if (&x.p).Cmp(&xt.p) != 0 {
		t.Error("128, no pass")
	}
	x = NewXn(0)
	for _, d := range r {
		o := d.Order()
		if o == 2 || o == 3 || o == 5 || o == 9 {
			x.Mul(x, d)
		}
	}
	x.PrintPoly()
	xt.Add(NewXn(256), NewXn(1))
	if (&x.p).Cmp(&xt.p) != 0 {
		t.Error("256, no pass")
	}
}

func TestFactorize(t *testing.T) {
	xt := NewXn(0)
	xt.Add(NewXn(128), NewXn(1))
	r := xt.Factorize()
	x := NewXn(0)
	for inx, p := range r {
		p.Println(fmt.Sprintf("Factor(%2d): ", inx))
		x.Mul(x, p)
	}
	Primes[len(Primes)-1].Println("Last Prime:")
	if (&x.p).Cmp(&xt.p) != 0 {
		t.Error("256 Fractorize, no pass")
	}
}
