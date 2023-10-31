package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewFloat(math.Pow(2.5, 25))
	b := big.NewFloat(math.Pow(1.9, 30))
	var c big.Float

	fmt.Printf("a = %.2f\n", a)
	fmt.Printf("b = %.2f\n", b)

	fmt.Printf("Sum %.2f\n", c.Add(a, b))
	fmt.Printf("Sub %.2f\n", c.Sub(a, b))
	fmt.Printf("Mul %.2f\n", c.Mul(a, b))
	fmt.Printf("div %.2f\n", c.Quo(a, b))
}
