package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func BinarySearch(slice []int, target int) int {

	left := 0
	right := len(slice) - 1
	var mid int

	for left < right {

		mid = (left + right) / 2

		if slice[mid] < target {
			left = mid + 1
		} else if slice[mid] >= target {
			right = mid
		}

	}

	if slice[right] == target {
		return right
	}

	return -1
}

func main() {

	for testCase := 0; testCase < 10; testCase++ {

		fmt.Printf("Test case #%d\n", testCase)

		sliceLen := rand.Intn(2) + 1
		slice := make([]int, sliceLen)
		for i := 0; i < len(slice); i++ {
			slice[i] = rand.Intn(5)
		}
		sort.Ints(slice)
		fmt.Printf("Slice: %v\n", slice)

		targetIndex := rand.Intn(sliceLen)
		fmt.Printf("Target: %d (index %d)\n", slice[targetIndex], targetIndex)

		result := BinarySearch(slice, slice[targetIndex])
		if result == -1 {
			fmt.Println("Result is -1, test failed, continuing")
		}

		fmt.Printf("Result: %d (index %d)\n", slice[result], result)

		fmt.Printf("Test result: %t\n", slice[result] == slice[targetIndex])

		fmt.Printf("\n%s\n\n", strings.Repeat("*", 40))
	}

}
