package main

import (
	"fmt"
	"time"
)

func SleepV1(ns int) {
	end := time.Now().Add(time.Duration(ns))
	for time.Now().Before(end) {
	}
}

func main() {
	start := time.Now()
	SleepV1(1000000000)
	fmt.Println(time.Now().Sub(start))
}
