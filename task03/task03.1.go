package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8, 10}

	c := make(chan int)

	for _, num := range nums {
		go func(num int) { c <- num * num }(num)
	}

	var sum int
	for range nums {
		sum += <-c
	}

	close(c)

	fmt.Println(sum)
}
