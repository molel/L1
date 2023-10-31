package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = i
	}

	firstChan := make(chan int, 2)
	secondChan := make(chan int, 2)

	go func() {
		for v := range numbers {
			firstChan <- v
		}
		close(firstChan)
	}()

	go func() {
		for v := range firstChan {
			secondChan <- v * v
		}
		close(secondChan)
	}()

	for v := range secondChan {
		fmt.Println(v)
		//time.Sleep(time.Millisecond * 200)
	}
}
