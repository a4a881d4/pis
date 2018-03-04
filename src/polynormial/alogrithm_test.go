package polynormail

import (
	"testing"
	"math/big"
)

func TestNewXn( t *testing.T ) {
	x := NewXn(1)
	y := big.Int(x)
	(&x).Println("Println")
	(&x).PrintPoly()
	if big.NewInt(2).Cmp(&y) != 0 {
		t.Error("xn = 1, no pass")
	}
}