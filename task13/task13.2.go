package main

import "fmt"

func main() {
	a := 0
	b := 1

	fmt.Printf("a = %d and b = %d\n", a, b)

	a += b
	b = a - b
	a = a - b

	fmt.Printf("a = %d and b = %d\n", a, b)
}
