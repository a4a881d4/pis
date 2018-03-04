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

func (x poly) Println(s string) {
	y := big.Int(x)
	fmt.Println(s,(&y).String())
}