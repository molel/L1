package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	// для завершения горутины отслеживается закрытие канала
	go func() {
		for value := range c {
			fmt.Printf("Recieved: %d\n", value)
		}
		fmt.Println("Goroutine finished")
	}()

	for i := 0; i < 3; i++ {
		c <- i
	}

	close(c)

	time.Sleep(time.Second * 2)
}
