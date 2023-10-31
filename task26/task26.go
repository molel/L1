package main

import (
	"fmt"
)

func IsUnique(s string) bool {
	m := make(map[rune]struct{})
	for _, r := range s {
		// если символ находится в верхнем регистре, он переводится в нижний
		if r < 97 {
			r += 32
		}
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}

func main() {
	s := "abcdA"
	fmt.Println(IsUnique(s))
}
