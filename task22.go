package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigIntA := big.NewInt(1<<44 + 42)
	bigIntB := big.NewInt(1<<44 + 14)

	fmt.Println(bigIntA)
	fmt.Println(bigIntB)

	fmt.Println("addition is ", addition(bigIntA, bigIntB))
	fmt.Println("subtraction is ", subtraction(bigIntA, bigIntB))
	fmt.Println("division is ", division(bigIntA, bigIntB))
	fmt.Println("multiplication is ", multiplication(bigIntA, bigIntB))
}

func addition(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Add(a, b)
}

func subtraction(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Sub(a, b)
}

func division(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Div(a, b)
}

func multiplication(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(a, b)
}
