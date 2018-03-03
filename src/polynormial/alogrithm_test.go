package polynormail

import (
	"testing"
	"math/big"
)

func TestNewXn( t *testing.T ) {
	x := NewXn(1)
	if big.NewInt(2).Cmp(&x) != 0 {
		// x.Println()
		t.Error("xn = 1, no pass")
	}
}