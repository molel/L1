package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// данный способ переворачивает предложение по словам с сохранением всех знаков препинания
	// для этого используются два отдельных шаблона регулярных выражений - для слов и для всех остальных символов

	sentence := ".Под тихим, звенящим, серебряным дождем, который лился стремительно, как небесная река, они встретились на скамейке, где судьба, будто грозовая молния, но с ласковой улыбкой, свела их вместе, и, словно волшебный вечер, их сердца начали биться сильнее, искреннее, и вечно/"

	wordsPattern, _ := regexp.Compile(`[a-zA-Zа-яА-Я0-9]+`)
	delimitersPattern, _ := regexp.Compile(`[^a-zA-Zа-яА-Я0-9]+`)

	words := wordsPattern.FindAllString(sentence, -1)
	delimiters := delimitersPattern.FindAllString(sentence, -1)

	if len(delimiters) < len(words) {
		delimiters = append(delimiters, "")
	}
	if len(delimiters) == len(words) {
		if wordsPattern.Match([]byte(string([]rune(sentence)[0]))) {
			delimiters = append(make([]string, 1, len(delimiters)), delimiters...)
		} else {
			delimiters = append(delimiters, "")
		}
	}

	var reversedString strings.Builder
	for i := 0; i < len(words); i++ {
		reversedString.WriteString(delimiters[i])
		reversedString.WriteString(words[len(words)-1-i])
	}
	reversedString.WriteString(delimiters[len(delimiters)-1])

	fmt.Println(reversedString.String())
}
