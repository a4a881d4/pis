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

func (p *PMatrix) NewMatrix() *PMatrix {
	A := make([]int64, p.M*p.N)
	for i, k := range p.A {
		A[i] = k
	}
	return &PMatrix{M: p.M, N: p.N, A: A, p: p.p}
}

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

func (p *PMatrix) FindMajor(n int) (int, int) {
	for i := n; i < p.M; i++ {
		for j := n; j < p.N; j++ {
			if p.A[j*p.M+i] != 0 {
				return i, j
			}
		}
	}
	return p.M, p.N
}
func (g *PMatrix) Set(i, j int, v int64) {
	g.A[j*g.M+i] = v
}
func (g *PMatrix) Get(i, j int) int64 {
	return g.A[j*g.M+i]
}
func (g *PMatrix) Guass() {
	order := make([]int, g.M)
	for n := 0; n < g.N; n++ {
		g.PrintMatrix("A")
		i, j := g.FindMajor(n)
		g.SelectMajor(i, j, n, order)
		g.PrintMatrix("FM")
		in := g.p.Inv(g.Get(n, n))
		fmt.Println("in", in)
		for i := n; i < g.M; i++ {
			g.A[n*g.M+i] = g.p.Mul(g.A[n*g.M+i], in)
		}
		g.PrintMatrix("GYH")
		for j := 0; j < g.N; j++ {
			if j != n {
				for i := n + 1; i < g.M; i++ {
					g.A[j*g.M+i] = g.p.Add(g.A[j*g.M+i], g.p.Mul(g.A[j*g.M+n], g.A[n*g.M+i]))
				}
				g.A[j*g.M+n] = 0
			}
			g.PrintMatrix("X")
		}
	}
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
				r.Set(i, j, A.p.Add(r.Get(i, j), A.p.Mul(A.Get(i, k), x.Get(k, j))))
			}
		}
	}
	return r
}

func (A *PMatrix) Part(j, i, m, n int) *PMatrix {
	d := make([]int, m*n)
	r := A.p.NewMatrix(m, n, d)
	for ii := 0; ii < n; ii++ {
		for jj := 0; jj < m; jj++ {
			r.Set(jj, ii, A.Get(j+jj, i+ii))
		}
	}
	return r
}
