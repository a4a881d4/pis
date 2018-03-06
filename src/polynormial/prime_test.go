package polynormal

import (
	"fmt"
	_ "math/big"
	"testing"
)

func TestP258(t *testing.T) {
	xt := NewXn(0)
	xt.Add(NewXn(258), NewXn(1))
	x := NewXn(0)
	for _, p := range P258 {
		x.Mul(x, p.ToPoly())
	}
	a := NewXn(0)
	x.Mul(x, a.Add(NewXn(2), NewXn(1)))
	x.Println("Println")
	x.PrintPoly()
	if (&xt.p).Cmp(&x.p) != 0 {
		t.Error("P258, no pass")
	}
}

func TestP258Mul(t *testing.T) {
	prime := P258[1]
	prime.Dump()
	a := NewRand(int(prime.order) - 1)
	b := NewRand(int(prime.order) - 1)
	a.Println("a")
	b.Println("b")
	r := prime.Mul((&a.p).Int64(), (&b.p).Int64())
	fmt.Printf("r: 0x%x\n", r)
	v := NewXn(0)
	v.Mul(a, b)
	v.Println("v")
	v.DivRem(prime.ToPoly())
	v.Println("v")
	if r != (&v.p).Int64() {
		t.Error("P258 Mul, no pass")
	}
}

func TestP128Mul(t *testing.T) {
	prime := P128[0]
	prime.Dump()
	a := NewRand(int(prime.order) - 1)
	b := NewRand(int(prime.order) - 1)
	a.Println("a")
	b.Println("b")
	r := prime.Mul((&a.p).Int64(), (&b.p).Int64())
	fmt.Printf("r: 0x%x\n", r)
	v := NewXn(0)
	v.Mul(a, b)
	v.Println("v")
	v.DivRem(prime.ToPoly())
	v.Println("v")
	if r != (&v.p).Int64() {
		t.Error("P128 Mul, no pass")
	}
}
