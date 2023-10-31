package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func QuickSort(slice []int, left, right int) {
	if left >= right {
		return
	}

	start, end := left, right
	mid := (left + right) / 2
	pivot := slice[mid]

	for left <= right {

		for left <= right && slice[left] < pivot {
			left++
		}

		for left <= right && slice[right] > pivot {
			right--
		}

		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}

	}

	QuickSort(slice, start, right)
	QuickSort(slice, left, end)
}

func main() {
	var sliceLen int
	for testCase := 0; testCase < 10; testCase++ {
		fmt.Printf("Test case #%d\n", testCase)

		sliceLen = rand.Intn(10)

		slice := make([]int, sliceLen)
		for i := 0; i < len(slice); i++ {
			slice[i] = rand.Intn(20)
		}
		fmt.Printf("Unsorted slice: %v\n", slice)

		QuickSort(slice, 0, sliceLen-1)
		fmt.Printf("Sorted slice: %v\n", slice)

		fmt.Println(strings.Repeat("*", 40))
	}
}
