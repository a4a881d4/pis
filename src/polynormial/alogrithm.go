package polynormail

import (
	"fmt"
	"math/big"
)

type poly big.Int

func NewXn(n int) poly {
	r := big.NewInt(1)
	for i:=0;i<n;i++ {
		r.Mul(r,big.NewInt(2))
	}
	return poly(*r)
}

func (x poly) Println() {
	fmt.Println(x)
}