package polynormal

import (
	"fmt"
	"math/big"
)

type Prime struct {
	poly  int64
	order int64
	size  int64
	idx   []int64
	power []int64
}

func NewPrime(a int64) *Prime {
	r := &Prime{poly: a}
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
	r := NewPrime(x.poly)
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
	x.Check()
	order := int64(x.Order())
	size := (int64(1) << uint(order-1))
	// div := make([]int64, size)
	// rem := make([]int64, size)
	idx := make([]int64, size)
	power := make([]int64, size)
	a := int64(1)
	for i := int64(0); i < size-1; i++ {
		power[i] = a
		idx[a] = i
		// fmt.Println("idx", idx[a], a, i)
		a <<= 1
		if (a>>uint(order-1))&1 == 1 {
			a ^= x.poly
		}
	}
	power[size-1] = 0
	idx[0] = size - 1
	x.idx = idx
	x.power = power
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
	fmt.Println("a", a, "b", b, x.idx[a], x.idx[b], (x.idx[a]+x.idx[b])%(x.size-1), x.power[(x.idx[a]+x.idx[b])%(x.size-1)])
	return x.power[(x.idx[a]+x.idx[b])%(x.size-1)]
}

func (x *Prime) Div(a, b int64) int64 {
	if a <= 0 || a >= (x.size-1) {
		return 0
	}
	if b == 0 || b >= (x.size-1) {
		return 0
	}

	return x.power[(x.idx[a]+(x.size-1)-x.idx[b])%(x.size-1)]
}

func (x *Prime) Inv(a int64) int64 {
	if a <= 0 || a >= (x.size-1) {
		return 0
	}
	return x.power[(x.size-1)-x.idx[a]]
}
func (x *Prime) Dump() {
	fmt.Println("Dump PrimPoly")
	fmt.Println("order", x.order)
	fmt.Println("size", x.size)
	for i := 0; i < int(x.size); i++ {
		fmt.Printf("%06x-%06x ", x.power[i], x.idx[i])
		if x.idx[i] != int64(0) && i > 32 {
			fmt.Println("exit")
			fmt.Println(i)
			break
		}
	}
	fmt.Println("")
}

type PolyBase struct {
	basisPoly []*Prime
}

func NewPolyBase(b []*Prime) *PolyBase {
	return &PolyBase{basisPoly: b}
}

func (b *PolyBase) Project(x *Poly) []int64 {
	r := make([]int64, len(b.basisPoly))
	for j := 0; j < x.Order(); j++ {
		if (&x.p).Bit(j) == 1 {
			for i, p := range b.basisPoly {
				r[i] ^= p.power[j]
			}
		}
	}
	return r
}
