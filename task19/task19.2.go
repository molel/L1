package main

import "fmt"

func main() {
	str := "Hello, 世界"
	var reversedStr string
	for _, v := range str {
		reversedStr = string(v) + reversedStr
	}
	fmt.Println(reversedStr)
}
