package main

import "fmt"

func main() {
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = i
	}

	fmt.Printf("Before deleting 3rd element: %v\n", nums)

	i := 3
	nums = append(nums[:i], nums[i+1:]...)

	fmt.Printf("After deleting 3rd element: %v\n", nums)
}
