package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterV2 struct {
	Count *uint32
}

func NewCounterV2() *CounterV2 {
	return &CounterV2{Count: new(uint32)}
}

// другой способ корректно увеличить счетчик в конкрутной среде - использовать атомарные операции
func (c CounterV2) Increment() {
	atomic.AddUint32(c.Count, 1)
}

func main() {
	counter := NewCounterV2()
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 100; i++ {
				//time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
				counter.Increment()
			}
		}()
	}
	wg.Wait()

	fmt.Println(*counter.Count)
}
