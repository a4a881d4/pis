package polynormal

import (
	"fmt"
	"math/big"
	"math/rand"
)

const PrimeSize = 1000

type Poly struct {
	p big.Int
}

var polyRand *rand.Rand
var Primes []*Poly

func init() {
	polyRand = rand.New(rand.NewSource(1))
	Primes = FindPrime(18)
}

func NewXn(n int) *Poly {
	if n < 0 {
		fmt.Println("Xn: must n >= 0", n)
	}
	r := big.NewInt(1)
	for i := 0; i < n; i++ {
		r.Mul(r, big.NewInt(2))
	}
	return &Poly{p: *r}
}

func NewPolyInt64(n int64) *Poly {
	r := big.NewInt(n)
	return &Poly{p: *r}
}

func NewRand(n int) *Poly {
	m := NewXn(n).p
	v := big.NewInt(0)
	v.Rand(polyRand, &m)
	return &Poly{p: *v}
}

func (x *Poly) Println(s string) {
	fmt.Println(s, (&x.p).Text(16))
}

func (x *Poly) Order() int {
	y := x.p
	i := 0
	r := big.NewInt(1)
	for {
		if r.Cmp(&y) <= 0 {
			i += 1
			r.Mul(r, big.NewInt(2))
		} else {
			break
		}

	}
	return i
}

func (x *Poly) PrintPoly() {
	o := x.Order()
	y := &(x.p)
	for i, j := o, 0; i > 0; i-- {
		if y.Bit(i) != 0 {
			fmt.Printf("X[%03d]", i)
			j++
			if j == 8 {
				j = 0
				fmt.Printf("\n")
			}
		}
	}
	if y.Bit(0) != 0 && o != 0 {
		fmt.Printf("1")
	}
	fmt.Println()
}

func (y *Poly) NewPoly() *Poly {
	x := big.NewInt(0)
	z := y.p
	x.SetBytes((&z).Bytes())
	return &Poly{p: *x}
}

func (z *Poly) Add(x, y *Poly) *Poly {
	r := z.p.Xor(&x.p, &y.p)
	z.p = *r
	return z
}

func (z *Poly) Mul(x, y *Poly) *Poly {
	s := big.NewInt(0)
	o := y.Order()
	vx := x.NewPoly()
	ix, iy := vx.p, y.p
	px, py := &ix, &iy
	for i := 0; i <= o; i++ {
		if py.Bit(i) != 0 {
			s.Xor(s, px)
		}
		px.Lsh(px, 1)
	}
	z.p = *s
	return z
}

func (a *Poly) DivRem(b *Poly) *Poly {
	oa := a.Order()
	ob := b.Order()
	r := big.NewInt(0)
	d := big.NewInt(0)
	if oa < ob {
		return &Poly{p: *r}
	}
	or := uint(oa - ob)
	d.Lsh(&b.p, or)
	pa := &a.p
	for i := 0; i < int(or+1); i++ {
		r.Lsh(r, 1)
		if pa.Bit(oa-i-1) != 0 {
			pa.Xor(pa, d)
			r.Xor(r, big.NewInt(1))
		}
		d.Rsh(d, 1)
	}
	return &Poly{p: *r}
}

func FindPrime(n int) []*Poly {
	max := &(NewXn(n/2 + 1).p)
	r := make([]*Poly, PrimeSize)
	cnt := 1
	r[0] = &Poly{p: *big.NewInt(2)}
	i := int64(3)
	for i < (int64(1) << uint(n)) {
		find := 0
		for _, d := range r[:cnt] {
			s := &Poly{p: *big.NewInt(i)}
			s.DivRem(d)
			if (&s.p).Sign() == 0 {
				find = 1
				break
			}
			if max.Cmp(&d.p) < 0 {
				break
			}
		}
		if find == 0 {
			if cnt < PrimeSize {
				r[cnt] = &Poly{p: *big.NewInt(i)}
				cnt++
			} else {
				break
			}
		}
		i++
	}
	return r[:cnt]
}

func (a *Poly) Factorize() []*Poly {
	d := a.NewPoly()
	r := make([]*Poly, PrimeSize)
	cnt := 0
	for _, p := range Primes {
		for {
			s := d.NewPoly()
			m := s.DivRem(p)
			if (&s.p).Sign() == 0 {
				r[cnt] = p
				cnt++
				d = m
			} else {
				break
			}
		}
		if d.Order() < p.Order() {
			break
		}
	}
	if d.p.Sign() != 0 {
		r[cnt] = d
		cnt++
	}
	return r[:cnt]
}

func (x *Poly) Euclidean(y *Poly) (*Poly, *Poly, *Poly) {
	a := make([]*Poly, 0, PrimeSize)
	b := make([]*Poly, 0, PrimeSize)
	swap := false
	if x.Order() >= y.Order() {
		a = append(a, x.NewPoly())
		a = append(a, y.NewPoly())
	} else {
		a = append(a, y.NewPoly())
		a = append(a, x.NewPoly())
		swap = true
	}
	for i := 0; i < PrimeSize; i++ {
		a = append(a, a[i].NewPoly())
		v := a[i+2].DivRem(a[i+1])
		b = append(b, v)
		if a[i+2].p.Sign() == 0 {
			break
		}
	}
	r := a[len(a)-2].NewPoly()
	m := NewXn(0)
	n := b[len(b)-2].NewPoly()
	for i := len(b) - 3; i >= 0; i-- {
		t := n.NewPoly()
		// fmt.Println(i)
		// m.Println("m")
		// n.Println("n")
		// fmt.Println("--------------")
		m, n = n, t.Mul(t, b[i]).Add(t, m)
	}
	if swap {
		return n, m, r
	} else {
		return m, n, r
	}
}
