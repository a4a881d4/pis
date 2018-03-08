package polynormal

import (
	"math/big"
)

var poly256 = [...]int64{
	0x11b,
	0x11d,
	0x12b,
	0x12d,
	0x139,
	0x13f,
	0x14d,
	0x15f,
	0x163,
	0x165,
	0x169,
	0x171,
	0x177,
	0x17b,
	0x187,
	0x18b,
	0x18d,
	0x19f,
	0x1a3,
	0x1a9,
	0x1b1,
	0x1bd,
	0x1c3,
	0x1cf,
	0x1d7,
	0x1dd,
	0x1e7,
	0x1f3,
	0x1f5,
	0x1f9}

var P256 []*Prime
var P256Base *PolyBase

func init() {
	P256 = make([]*Prime, len(poly256))
	for i, p := range poly256 {
		P256[i] = NewPrime(p, true)
	}
	P256Base = NewPolyBase(P256)
}

type PSign struct {
	Gen *Poly
	S   []*Poly
}

func toInt(message [][32]byte, m, n int) []int {
	r := make([]int, m*n)
	for i := 0; i < n; i++ {
		v := message[i]
		for j := 0; j < m; j++ {
			r[i*m+j] = int(v[j]) & 0xff
		}
	}
	return r
}

func NewSign256(message [][32]byte) *PSign {
	n := len(message)
	m := n + 1
	g, v := P256Base.Solution(m, n, 16, toInt(message, m, n))
	return &PSign{Gen: g, S: v}
}

func (c *PSign) Check256(message [32]byte) bool {
	n := len(c.S)
	m := n + 1
	ms := [...][32]byte{message}
	return c.Gen.CheckSolution(0xff, c.S, toInt(ms, m, 1))
}
