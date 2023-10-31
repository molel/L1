package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"time"
)

func Work(ctx context.Context, readChan chan interface{}, number int) {
	for {
		select {
		case data := <-readChan:
			fmt.Printf("Worker #%d recieved data: %v\n", number, data)
			time.Sleep(time.Nanosecond * 200)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	readChan := make(chan interface{})
	defer close(readChan)

	// ctrl + C передает сигнал os.Interrupt
	// поэтому используется контекст, который завершается при получении указанного сигнала
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var numberOfWorkers int
	fmt.Print("Choose number of workers: ")
	_, err := fmt.Scan(&numberOfWorkers)
	if err != nil {
		fmt.Printf("Cannot scan the value: %s", err)
		return
	}

	for i := 0; i < numberOfWorkers; i++ {
		go Work(ctx, readChan, i)
	}

	go func() {
		for i := 0; i < math.MaxInt; i++ {
			readChan <- i
		}
	}()

	<-ctx.Done()
}
