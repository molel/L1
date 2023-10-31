package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	nums := []int32{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	var sum int32
	for _, num := range nums {
		wg.Add(1)
		go func(num int32) {
			defer wg.Done()
			atomic.AddInt32(&sum, num*num)
		}(num)
	}

	wg.Wait()

	fmt.Println(sum)
}
