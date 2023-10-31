package main

import (
	"fmt"
	"sync"
)

func CalculateSquares(nums []int, wg *sync.WaitGroup) {
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(num)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	// передача указателя на WaitGroup для избежания дедлока
	CalculateSquares([]int{2, 4, 6, 8, 10}, wg)
	wg.Wait()
}
