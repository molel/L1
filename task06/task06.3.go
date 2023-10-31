package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	readChan := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())

	// для завершения горутины используется контекст с фукнцией завершения
	go func(ctx context.Context) {
		for {
			select {
			case readChan <- "something":
				time.Sleep(time.Millisecond * 400)
			case <-ctx.Done():
				fmt.Println("Goroutine finished")
				close(readChan)
				return
			}
		}
	}(ctx)

	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	for value := range readChan {
		fmt.Printf("Recieved: %s\n", value)
	}
}
