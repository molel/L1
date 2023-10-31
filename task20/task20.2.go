package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// данный способ переворачивает предложение по словам без сохраненя всех знаков препинания

	sentence := ".Под тихим, звенящим, серебряным дождем, который лился стремительно, как небесная река, они встретились на скамейке, где судьба, будто грозовая молния, но с ласковой улыбкой, свела их вместе, и, словно волшебный вечер, их сердца начали биться сильнее, искреннее, и вечно/"

	wordsPattern, _ := regexp.Compile(`[a-zA-Zа-яА-Я0-9]+`)

	words := wordsPattern.FindAllString(sentence, -1)

	var reversedString strings.Builder
	for i := len(words) - 1; i >= 0; i-- {
		reversedString.WriteString(words[i])
		reversedString.WriteString(" ")
	}

	fmt.Println(reversedString.String())
}
