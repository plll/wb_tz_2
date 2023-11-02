package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func div(a, b *big.Int) *big.Int {
	z := new(big.Int)
	return z.Div(a, b)
}

func add(a, b *big.Int) *big.Int {
	z := new(big.Int)
	return z.Add(a, b)
}

func mul(a, b *big.Int) *big.Int {
	z := new(big.Int)
	return z.Mul(a, b)
}

func sub(a, b *big.Int) *big.Int {
	z := new(big.Int)
	return z.Sub(a, b)
}

func main() {
	a := new(big.Int)
	a.SetString("24124124456789028482358238582358238582385283582385283582385235823858212453647858963541231215151274743", 10)
	b := new(big.Int)
	b.SetString("23456789123418128481248184182481248184182481248", 10)
	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(div(a, b))
	fmt.Println(mul(a, b))
}
