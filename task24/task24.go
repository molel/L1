package main

import (
	"L1/task24/point"
	"fmt"
)

func main() {
	a := point.NewPoint(0, 3)
	b := point.NewPoint(4, 0)

	fmt.Println(a.DistanceTo(b))
}
