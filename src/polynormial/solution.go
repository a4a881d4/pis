package polynormal

import (
	_ "fmt"
)

func (b *PolyBase) Solution(m, n, maxb int, c []int) (*Poly, []*Poly) {
	basis := make([]*Prime, 0, len(b.basisPoly))
	y := make(map[*Prime]*PMatrix)
	for _, p := range b.basisPoly {
		A := p.NewMatrix(m, n, c)
		Y, invable := A.Guass(false)
		if invable {
			basis = append(basis, p)
			y[p] = Y
			if len(basis) >= maxb {
				break
			}
		}
	}
	// for k, v := range y {
	// 	k.ToPoly().Println("k")
	// 	for _, xi := range v.A {
	// 		fmt.Printf("%03x ", xi)
	// 	}
	// 	fmt.Println("")
	// }
	newBase := NewPolyBase(basis)
	r := make([]*Poly, n)
	for i := 0; i < n; i++ {
		x := make([]int64, len(basis))
		for k, p := range basis {
			x[k] = y[p].A[i]
		}
		r[i] = newBase.Synthesize(x)
		// r[i].Println("ri")
	}
	return newBase.products, r
}

func (products *Poly) CheckSolution(s int, x []*Poly, c []int) bool {
	sum := NewPolyInt64(0)
	for i, p := range x {
		if c[i] != s {
			xn := NewXn(c[i])
			sum.Add(sum, xn.Mul(xn, p))
			// sum.Println("sum")
		}
	}
	sum.Add(sum, NewXn(c[len(c)-1]))
	// r := NewXn(0)
	// sum.DivRem(r.Add(r, NewXn(s)))
	// sum.Println("sum")
	// products.Println("Products")
	sum.DivRem(products)
	if sum.p.Sign() == 0 {
		return true
	} else {
		return false
	}
}
