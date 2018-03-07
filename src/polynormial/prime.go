package polynormal

import (
	"fmt"
	"math/big"
)

type Prime struct {
	poly      int64
	order     int64
	size      int64
	indexable bool
	index     []int64
	power     []int64
}

func NewPrime(a int64, gen bool) *Prime {
	r := &Prime{poly: a, indexable: gen}
	r.genTable()
	return r
}

func (x *Prime) ToPoly() *Poly {
	v := big.NewInt(x.poly)
	return &Poly{p: *v}
}

func (x *Prime) Order() int {
	return x.ToPoly().Order()
}

func (x *Prime) NewPrime() *Prime {
	r := NewPrime(x.poly, x.indexable)
	return r
}

func (x *Prime) Check() {
	vx := x.ToPoly()
	ps := vx.Factorize()
	for _, p := range ps {
		p.Println("x: ")
	}
}
func (x *Prime) genTable() {
	order := int64(x.Order())
	size := (int64(1) << uint(order-1))
	if x.indexable {
		index := make([]int64, size)
		power := make([]int64, size)
		a := int64(1)
		for i := int64(0); i < size-1; i++ {
			power[i] = a
			index[a] = i
			a <<= 1
			if (a>>uint(order-1))&1 == 1 {
				a ^= x.poly
			}
		}
		power[size-1] = 0
		index[0] = size - 1
		x.index = index
		x.power = power
	}
	x.order = order
	x.size = size
}

func (x *Prime) Add(a, b int64) int64 {
	return a ^ b
}

func (x *Prime) Mul(a, b int64) int64 {
	if a <= 0 || a >= (x.size-1) {
		return 0
	}
	if b == 0 || b >= (x.size-1) {
		return 0
	}
	if x.indexable {
		return x.power[(x.index[a]+x.index[b])%(x.size-1)]
	} else {
		v := NewXn(0)
		v.Mul(NewPolyInt64(a), NewPolyInt64(b))
		v.DivRem(x.ToPoly())
		return v.p.Int64()
	}
}

func (x *Prime) Div(a, b int64) int64 {
	if a <= 0 || a >= (x.size-1) {
		return 0
	}
	if b == 0 || b >= (x.size-1) {
		return 0
	}
	if x.indexable {
		return x.power[(x.index[a]+(x.size-1)-x.index[b])%(x.size-1)]
	} else {
		return x.Mul(a, x.Inv(b))
	}
}

func (x *Prime) Inv(a int64) int64 {
	if a <= 0 || a >= (x.size-1) {
		return 0
	}
	if x.indexable {
		if x.index[a] != 0 {
			return x.power[(x.size-1)-x.index[a]]
		} else {
			return 1
		}
	} else {
		_, n, _ := x.ToPoly().Euclidean(NewPolyInt64(a))
		return n.p.Int64()
	}
}
func (x *Prime) Dump() {
	fmt.Println("Dump PrimPoly")
	fmt.Println("order", x.order)
	fmt.Println("size", x.size)
	if x.indexable {
		for i := 0; i < int(x.size); i++ {
			fmt.Printf("%06x-%06x ", x.power[i], x.index[i])
			if i > 32 {
				fmt.Println("exit")
				fmt.Println(i)
				break
			}
		}
	}
	fmt.Println("")
}

type PolyBase struct {
	basisPoly []*Prime
	products  *Poly
	bi        []*Poly
}

func NewPolyBase(b []*Prime) *PolyBase {
	products := NewXn(0)
	for _, p := range b {
		products.Mul(products, p.ToPoly())
	}
	r := make([]*Poly, len(b))
	for i, p := range b {
		ps := products.NewPoly()
		r[i] = ps.DivRem(p.ToPoly())
		pi := r[i].NewPoly()
		pi.DivRem(p.ToPoly())
		r[i].Mul(r[i], NewPolyInt64(p.Inv(pi.p.Int64())))
		r[i].DivRem(products)
	}
	return &PolyBase{basisPoly: b, products: products, bi: r}
}

func (b *PolyBase) Analysis(x *Poly) []int64 {
	r := make([]int64, len(b.basisPoly))
	for j := 0; j < x.Order(); j++ {
		if x.p.Bit(j) == 1 {
			for i, p := range b.basisPoly {
				if p.indexable {
					r[i] ^= p.power[j]
				}
			}
		}
	}
	for i, p := range b.basisPoly {
		if !p.indexable {
			vx := x.NewPoly()
			vx.DivRem(p.ToPoly())
			r[i] = vx.p.Int64()
		}
	}
	return r
}

func (b *PolyBase) Synthesize(f []int64) *Poly {
	r := NewPolyInt64(0)
	for i, fi := range f {
		x := NewPolyInt64(fi)
		r.Add(r, x.Mul(x, b.bi[i]))
	}
	r.DivRem(b.products)
	return r
}
