package polynormail

import (
	"testing"
	"math/big"
)

func TestNewXn( t *testing.T ) {
	if big.NewInt(2) != NewXn(1) {
		t.Error("xn = 1, no pass")
	}
}