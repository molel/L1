package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[float64][]float64)

	for _, number := range numbers {
		// вычисление нижней границы диапазона числа
		key := math.Floor(number/10) * 10

		groups[key] = append(groups[key], number)
	}

	for k, v := range groups {
		fmt.Printf("%.0f: %v\n", k, v)
	}
}
