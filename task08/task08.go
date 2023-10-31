package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func SetBit(a int64, bitNumber int, bitValue int) int64 {
	switch bitValue {
	case 0:
		a &^= 1 << bitNumber
	case 1:
		a |= 1 << bitNumber
	}

	return a
}

func main() {
	bitNumber := rand.Intn(64)
	mask := []byte(strings.Repeat(" ", 64))
	mask[63-bitNumber] = '1'

	fmt.Printf("%064b\n", 0)
	fmt.Println(string(mask))
	fmt.Printf("%064b\n", SetBit(0, bitNumber, 1))

	fmt.Printf("\n%s\n\n", strings.Repeat("=", 64))

	fmt.Printf("%064b\n", math.MaxInt)
	bitNumber = rand.Intn(64)
	mask = []byte(strings.Repeat(" ", 64))
	mask[63-bitNumber] = '0'
	fmt.Println(string(mask))
	fmt.Printf("%064b\n", SetBit(math.MaxInt, bitNumber, 0))
}
