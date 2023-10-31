package main

import "fmt"

func intersection(setA, setB map[int]struct{}) map[int]struct{} {
	set := make(map[int]struct{})

	for i, _ := range setA {
		if _, ok := setB[i]; ok {
			set[i] = struct{}{}
		}
	}

	for i, _ := range setB {
		if _, ok := setA[i]; ok {
			set[i] = struct{}{}
		}
	}

	return set
}

func main() {
	numsA := []int{0, 1, 2, 3, 4, 5, 6}
	numsB := []int{4, 5, 6, 7, 8, 9, 10}

	// в качестве значений хэш таблицы используются пустые структуры, так как они занимают минимум памяти - 0 байт
	setA := make(map[int]struct{})
	setB := make(map[int]struct{})

	for _, num := range numsA {
		setA[num] = struct{}{}
	}

	for _, num := range numsB {
		setB[num] = struct{}{}
	}

	fmt.Println(intersection(setA, setB))
}
