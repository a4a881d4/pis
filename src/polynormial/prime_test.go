package polynormal

import (
	_ "fmt"
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
	if xt.p.Cmp(&x.p) != 0 {
		t.Error("P258, no pass")
	}
}

func TestP258Mul(t *testing.T) {
	prime := P258[1]
	a := NewRand(int(prime.order) - 1)
	b := NewRand(int(prime.order) - 1)
	r := prime.Mul(a.p.Int64(), b.p.Int64())
	v := NewXn(0)
	v.Mul(a, b)
	v.DivRem(prime.ToPoly())
	if r != v.p.Int64() {
		t.Error("P258 Mul, no pass")
	}
}

func TestP128Mul(t *testing.T) {
	prime := P128[0]
	a := NewRand(int(prime.order) - 1)
	b := NewRand(int(prime.order) - 1)
	r := prime.Mul((&a.p).Int64(), (&b.p).Int64())
	v := NewXn(0)
	v.Mul(a, b)
	v.DivRem(prime.ToPoly())
	if r != v.p.Int64() {
		t.Error("P128 Mul, no pass")
	}
}

func TestP128Euclidean(t *testing.T) {
	prime := P128[0]
	a := NewRand(int(prime.order) - 1)
	v2 := a.NewPoly()
	m, n, _ := prime.ToPoly().Euclidean(a)
	if prime.Inv(n.p.Int64()) != a.p.Int64() {
		t.Error("P128 Euclidean Inv, no pass")
	}
	v1 := prime.ToPoly()
	s := v1.Mul(v1, m).Add(v1, v2.Mul(v2, n))
	if 1 != s.p.Int64() {
		t.Error("P128 Euclidean, no pass")
	}
}

func TestP258Euclidean(t *testing.T) {
	prime := P258[0]
	for i := 0; i < 256; i++ {
		a := NewRand(int(prime.order) - 1)
		v2 := a.NewPoly()
		m, n, _ := prime.ToPoly().Euclidean(a)
		v1 := prime.ToPoly()
		s := v1.Mul(v1, m).Add(v1, v2.Mul(v2, n))
		if 1 != s.p.Int64() {
			t.Error("P258 Euclidean, no pass")
		}
	}
}
