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

func TestP2048Inv(t *testing.T) {
	prime := P2048[0]
	a := NewRand(int(prime.order) - 1)
	c := prime.Inv(a.p.Int64())
	if prime.Mul(a.p.Int64(), c) != 1 {
		t.Error("P2048 Inv, no pass")
	}
}

func TestP128AS(t *testing.T) {
	base := NewPolyBase(P128)
	x := NewRand(126)
	a := base.Analysis(x)
	for i, ai := range a {
		ax := x.NewPoly()
		ax.DivRem(P128[i].ToPoly())
		if ax.p.Int64() != ai {
			t.Error("P128 Analysis, no pass")
		}
	}
	r := base.Synthesize(a)
	if r.p.Cmp(&x.p) != 0 {
		r.Println("r")
		x.Println("x")
		t.Error("P128 Synthesize, no pass")
	}
}

func checkP2048Basis(n int, t *testing.T) {
	base := NewPolyBase(P2048)
	y := NewXn(0)
	for i := 0; i < len(P2048); i++ {
		y.Mul(y, P2048[i].ToPoly())
	}
	x := y.DivRem(P2048[n].ToPoly())
	a := base.Analysis(x)
	for i, ai := range a {
		ax := x.NewPoly()
		ax.DivRem(P2048[i].ToPoly())
		if ax.p.Int64() != ai {
			t.Error("P2048 Analysis, no pass")
		}
	}
	r := base.Synthesize(a)
	if r.p.Cmp(&x.p) != 0 {
		r.Println("r")
		x.Println("x")
		NewPolyInt64(a[n]).Println("ai")
		fmt.Println("n", n)
		t.Error("P2048 Synthesize, no pass")
	}
}

func TestP2048Basis(t *testing.T) {
	testcase := [...]int{43, 64, 102, 173}
	for _, n := range testcase {
		checkP2048Basis(n, t)
	}
}

func TestP2048AS(t *testing.T) {
	base := NewPolyBase(P2048)
	for n := 0; n < 1; n++ {
		x := NewRand(1930)
		a := base.Analysis(x)
		for i, ai := range a {
			ax := x.NewPoly()
			ax.DivRem(P2048[i].ToPoly())
			if ax.p.Int64() != ai {
				t.Error("P2048 Analysis, no pass")
			}
		}
		r := base.Synthesize(a)
		if r.p.Cmp(&x.p) != 0 {
			r.Println("r")
			x.Println("x")
			t.Error("P2048 Synthesize, no pass")
		}
	}
}
