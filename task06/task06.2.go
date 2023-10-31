package main

import (
	"fmt"
	"time"
)

func main() {
	readChan := make(chan string)
	stopChan := make(chan struct{})

	// для завершения горутины используется отдельный канал,
	// в который передается сигнал для завершения
	go func() {
		for {
			select {
			case readChan <- "something":
			case <-stopChan:
				fmt.Println("Goroutine finished")
				close(readChan)
				return
			}
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		time.Sleep(time.Second * 3)
		stopChan <- struct{}{}
	}()

	for value := range readChan {
		fmt.Printf("Recieved: %s\n", value)
	}
}
