package polynormal

import (
	"fmt"
)

type PMatrix struct {
	M, N int
	A    []int64
	p    *Prime
}

func (p *Prime) NewMatrix(m, n int, c []int) *PMatrix {
	A := make([]int64, m*n)
	for i, k := range c {
		if int64(k) == p.size-1 {
			A[i] = 0
		} else {
			if p.indexable {
				A[i] = p.power[k]
			} else {
				x := NewXn(k)
				x.DivRem(p.ToPoly())
				A[i] = x.p.Int64()
			}
		}
	}
	return &PMatrix{M: m, N: n, A: A, p: p}
}

func (p *Prime) NewMatrixInt64(m, n int, c []int64) *PMatrix {
	A := make([]int64, m*n)
	for i, k := range c {
		A[i] = k
	}
	return &PMatrix{M: m, N: n, A: A, p: p}
}

func (p *PMatrix) NewMatrix() *PMatrix {
	A := make([]int64, p.M*p.N)
	for i, k := range p.A {
		A[i] = k
	}
	return &PMatrix{M: p.M, N: p.N, A: A, p: p.p}
}

/*
j ------ M
i
|
|
N
*/
func (p *PMatrix) PrintMatrix(s string) {
	fmt.Println(s)
	for i := 0; i < p.N; i++ {
		for j := 0; j < p.M; j++ {
			fmt.Printf("%03x ", p.Get(j, i))
		}
		fmt.Println("")
	}
}

func (p *PMatrix) SelectMajor(i, j, n int, col []int) {
	if i >= p.N || n >= p.N {
		return
	}
	if j >= p.M || n >= p.M {
		return
	}
	if j != n {
		col[n], col[j] = col[j], col[n]
		for l := 0; l < p.N; l++ {
			p.A[l*p.M+n], p.A[l*p.M+j] = p.A[l*p.M+j], p.A[l*p.M+n]
		}
	}
	if i != n {
		for l := 0; l > p.M; l++ {
			p.A[i*p.M+l], p.A[n*p.M+l] = p.A[n*p.M+l], p.A[i*p.M+l]
		}
	}
}

func (p *PMatrix) FindMajor(n int) (int, int, bool) {
	for i := n; i < p.N; i++ {
		for j := n; j < p.N; j++ {
			if p.A[j*p.M+i] != 0 {
				return i, j, true
			}
		}
	}
	return p.M, p.N, false
}
func (g *PMatrix) Set(i, j int, v int64) {
	g.A[j*g.M+i] = v
}
func (g *PMatrix) Get(i, j int) int64 {
	return g.A[j*g.M+i]
}
func (g *PMatrix) Guass(v bool) (*PMatrix, bool) {
	order := make([]int, g.M)
	if g.M < g.N {
		return nil, false
	}
	for i, _ := range order {
		order[i] = i
	}
	for n := 0; n < g.N; n++ {
		i, j, invable := g.FindMajor(n)
		if !invable {
			return nil, false
		}
		g.SelectMajor(i, j, n, order)
		in := g.p.Inv(g.Get(n, n))
		if v {
			fmt.Printf("%03x inv %03x\n", g.Get(n, n), in)
		}
		for i := n; i < g.M; i++ {
			g.A[n*g.M+i] = g.p.Mul(g.A[n*g.M+i], in)
		}
		for j := 0; j < g.N; j++ {
			if j != n {
				for i := n + 1; i < g.M; i++ {
					g.A[j*g.M+i] = g.p.Add(g.A[j*g.M+i], g.p.Mul(g.A[j*g.M+n], g.A[n*g.M+i]))
				}
				g.A[j*g.M+n] = g.p.Add(g.A[j*g.M+n], g.p.Mul(g.A[j*g.M+n], g.A[n*g.M+n]))
			}
		}
	}
	r := g.p.NewMatrixInt64(g.M-g.N, g.N, make([]int64, (g.M-g.N)*g.N))
	for i := 0; i < g.N; i++ {
		for j := 0; j < g.M-g.N; j++ {
			r.Set(j, order[i], g.Get(g.N+j, i))
		}
	}
	return r, true
}

func (A *PMatrix) Mul(x *PMatrix) *PMatrix {
	if A.M != x.N {
		return nil
	}
	m := x.M
	n := A.N
	d := make([]int, m*n)
	r := A.p.NewMatrix(m, n, d)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			r.Set(i, j, 0)
			for k := 0; k < A.M; k++ {
				r.Set(i, j, A.p.Add(r.Get(i, j), A.p.Mul(A.Get(k, i), x.Get(j, k))))
			}

		}
	}
	return r
}

func (A *PMatrix) Part(i, j, m, n int) *PMatrix {
	d := make([]int, m*n)
	r := A.p.NewMatrix(m, n, d)
	for jj := 0; jj < n; jj++ {
		for ii := 0; ii < m; ii++ {
			r.Set(ii, jj, A.Get(i+ii, j+jj))
		}
	}
	return r
}

func (A *PMatrix) Equ(B *PMatrix) bool {
	if A.p.poly != B.p.poly {
		return false
	}
	if A.M != B.M || A.N != B.N {
		return false
	}
	for i, k := range A.A {
		if B.A[i] != k {
			return false
		}
	}
	return true
}
