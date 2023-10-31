package main

import "fmt"

func Set(strs []string) map[string]struct{} {
	// в качестве значений хэш таблицы используются пустые структуры, так как они занимают минимум памяти - 0 байт
	set := make(map[string]struct{})

	for _, str := range strs {
		set[str] = struct{}{}
	}

	return set
}

func main() {
	fmt.Println(Set([]string{"cat", "cat", "dog", "cat", "tree"}))
}
