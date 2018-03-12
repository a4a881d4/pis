package polynormal

import (
	_ "fmt"
)

func (b *PolyBase) Solution(m, n, maxb int, c []int) (*Poly, []*Poly) {
	basis := make([]*Prime, 0, len(b.basisPoly))
	y := make([]*PMatrix, 0, len(b.basisPoly))
	for _, p := range b.basisPoly {
		A := p.NewMatrix(m, n, c)
		Y, invable := A.Guass(false)
		if invable {
			basis = append(basis, p)
			y = append(y, Y)
			if len(basis) >= maxb {
				break
			}
		}
	}
	// for k, v := range y {
	// 	basis[k].ToPoly().Println("k")
	// 	for _, xi := range v.A {
	// 		fmt.Printf("%03x ", xi)
	// 	}
	// 	fmt.Println("")
	// }
	newBase := NewPolyBase(basis)
	r := make([]*Poly, n)
	for i := 0; i < n; i++ {
		x := make([]int64, len(basis))
		for k, _ := range basis {
			x[k] = y[k].A[i]
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
	if c[len(c)-1] != s {
		sum.Add(sum, NewXn(c[len(c)-1]))
	}
	// r := NewXn(0)
	// sum.DivRem(r.Add(r, NewXn(s)))
	// sum.Println("sum")
	// products.Println("Products")
	sum.DivRem(products)
	if sum.p.Sign() == 0 {
		return true
	} else {
		sum.Println("sum")
		for _, p := range x {
			p.Println("x")
		}
		return false
	}
}

func (b *PolyBase) NSolution(m, n, maxb int, c []int) ([]int, []*Prime) {
	basis := make([]*Prime, 0, len(b.basisPoly))
	y := make([]*PMatrix, 0, len(b.basisPoly))
	for _, p := range b.basisPoly {
		A := p.NewMatrix(m, n, c)
		// fmt.Println("k", k)
		// A.PrintMatrix("A")
		Y, invable := A.Guass(false)
		// Y.PrintMatrix("Y")
		if invable {
			basis = append(basis, p)
			y = append(y, Y)
			if len(basis) >= maxb {
				break
			}
		}
	}
	r := make([]int, n*len(basis))
	for i := 0; i < n; i++ {
		for k, p := range basis {
			r[i*len(basis)+k] = int(p.index[y[k].A[i]])
		}
	}
	return r, basis
}

func (p *Prime) NCheckSolution(m, n, pos, pk int, c, r []int) bool {
	B := p.NewMatrix(n, 1, c[0+m*pos:n+m*pos])
	// B.PrintMatrix("B")
	lbasis := len(r) / n
	C := p.NewMatrix(lbasis, n, r)
	// C.PrintMatrix("C")
	D := B.Mul(C)
	// D.PrintMatrix("D")
	// fmt.Printf("M %02x\n", p.power[c[n+m*pos]])
	if D.A[pk] == p.power[c[n+m*pos]] {
		return true
	} else {
		return false
	}
}
