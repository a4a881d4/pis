package polynormal

import (
	_ "fmt"
	"testing"
)

func TestP128Solution5X4(t *testing.T) {
	m := make([]int, 20)
	for i := 0; i < 20; i++ {
		m[i] = polyRand.Intn(127)
	}
	base := NewPolyBase(P128)
	b, x := base.Solution(5, 4, 10, m)
	for i := 0; i < 4; i++ {
		if !b.CheckSolution(0x7f, x, m[i*5:i*5+5]) {
			t.Error("P128 Solution, no pass")
		}
	}
}

func TestSign256(t *testing.T) {
	messageString := [...]string{
		"a4a881d4",
		"a4a881d4@163.com",
		"hello world",
		"proof in set",
		"Bitcoin uses peer-to-peer technology to operate with no central authority or banks",
		"managing transactions and the issuing of bitcoins is carried out collectively by the network",
		"Bitcoin is open-source",
		"its design is public",
		"nobody owns or controls Bitcoin and everyone can take part",
		"Through many of its unique properties",
		"Bitcoin allows exciting uses that could not be covered by any previous payment system",
		""
	}
}