package main

import (
	"fmt"
	"sync"
)

// корректная работа в конкурентной среде достигается с помощью мьютекса
type CounterV1 struct {
	sync.Mutex
	Count int
}

func NewCounterV1() *CounterV1 {
	return &CounterV1{}
}

func (c *CounterV1) Increment() {
	c.Lock()
	defer c.Unlock()

	c.Count++
}

func main() {
	counter := NewCounterV1()
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

	fmt.Println(counter.Count)
}
