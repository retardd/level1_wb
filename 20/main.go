package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	// split строки на слова
	words := strings.Fields(input)

	// реверс каждого слова
	for i := 0; i < len(words); i++ {
		words[i] = reverseString(words[i])
	}

	// join перевернутых слов в строку
	result := strings.Join(words, " ")

	return result
}

func reverseString(s string) string {
	var ReverseString string

	// строка в срез unicode-символов
	runes := []rune(s)
	tempLen := len(runes)

	for tempLen > 0 {
		tempLen--
		tempChar := runes[tempLen : tempLen+1]
		ReverseString = fmt.Sprintf("%s%s", ReverseString, string(tempChar))

	}
	return ReverseString
}

func main() {
	input := "snow dog sun — sun dog snow 123"
	result := reverseWords(input)

	fmt.Println("обработанная строка:", result)
}
