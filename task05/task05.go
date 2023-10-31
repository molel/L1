package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	N := 10

	// контекст, завершающийся по таймауту
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(N))
	defer cancel()

	c := make(chan int)
	defer close(c)

	go func(ctx context.Context, c chan int) {
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-c:
				fmt.Printf("Recieved data: %d\n", data)
			}
		}
	}(ctx, c)

	var i int
	for {
		select {
		case c <- i:
			i++
			time.Sleep(time.Second)
		case <-ctx.Done():
			return
		}
	}
}
