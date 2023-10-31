package main

import (
	"fmt"
)

func main() {
	str := "Hello, 世界"
	// для обработки символов юникода строка разбивается на руны
	// альтернативный способ разбития на руны - использование ключевого слова range
	runedStr := []rune(str)
	for i, j := 0, len(runedStr)-1; i < j; i, j = i+1, j-1 {
		runedStr[i], runedStr[j] = runedStr[j], runedStr[i]
	}
	fmt.Println(string(runedStr))
}
