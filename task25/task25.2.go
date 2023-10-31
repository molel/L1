package main

import (
	"fmt"
	"time"
)

func SleepV2(ns int) {
	<-time.After(time.Duration(ns))
}

func main() {
	start := time.Now()
	SleepV2(1000000000)
	fmt.Println(time.Now().Sub(start))
}
