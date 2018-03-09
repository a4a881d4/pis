package polynormal

import (
	"crypto/sha256"
	"fmt"
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

func TestP256Solution5X4(t *testing.T) {
	m := make([]int, 20)
	for i := 0; i < 20; i++ {
		m[i] = polyRand.Intn(255)
	}
	base := NewPolyBase(P256)
	b, x := base.Solution(5, 4, 10, m)
	for i := 0; i < 4; i++ {
		if !b.CheckSolution(0xff, x, m[i*5:i*5+5]) {
			t.Error("P256 Solution, no pass")
		}
	}
}

var messageString = [...]string{
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
	"end"}

func TestSign256(t *testing.T) {
	// messageString := [...]string{
	// 	"a4a881d4",
	// 	"a4a881d4@163.com",
	// 	"hello world",
	// 	"proof in set",
	// 	"Bitcoin uses peer-to-peer technology to operate with no central authority or banks",
	// 	"managing transactions and the issuing of bitcoins is carried out collectively by the network",
	// 	"Bitcoin is open-source",
	// 	"its design is public",
	// 	"nobody owns or controls Bitcoin and everyone can take part",
	// 	"Through many of its unique properties",
	// 	"Bitcoin allows exciting uses that could not be covered by any previous payment system",
	// 	"end"}
	message := make([][32]byte, len(messageString), 1024+32)
	for k, s := range messageString {
		message[k] = sha256.Sum256([]byte(s))
	}
	for i := 0; i < 32; i++ {
		for k, s := range message {
			message[k] = sha256.Sum256(s[:])
		}
		sign := NewSign256(message)
		for k, m := range message {
			if !sign.Check256(m) {
				fmt.Printf("%x\n", m)
				Check256(message)
				t.Error(messageString[k], " Check, no pass")
			}
			if sign.Gen.Order() != 113 {
				t.Error("miss order")
			}
		}
	}
	for k, s := range messageString {
		message[k] = sha256.Sum256([]byte(s))
	}
	for i := len(message); i < 31; i++ {
		for k, s := range message {
			message[k] = sha256.Sum256(s[:])
		}
		sign := NewSign256(message)
		for k, m := range message {
			if !sign.Check256(m) {
				fmt.Printf("%x\n", m)
				Check256(message)
				t.Error(k, " Check, no pass")
			}
			if sign.Gen.Order() != 113 {
				fmt.Printf("%d: %x\n", sign.Gen.Order(), m)
				t.Error(k, "miss order")
			}
		}
		message = append(message, sha256.Sum256([]byte(messageString[0])))
	}
}

func TestNSolution(t *testing.T) {
	message := make([][32]byte, len(messageString), 1024+32)
	for k, s := range messageString {
		message[k] = sha256.Sum256([]byte(s))
	}
	n := len(message)
	m := n + 1
	P256Base.NSolution(m, n, 14, toInt(message, m, n))
}
