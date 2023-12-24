package main

import (
	"fmt"
	"strings"
)

func isCharsUnique(s string) bool {
	// мапа с unicode-символами
	cch := make(map[rune]struct{})

	// lower строки
	lString := strings.ToLower(s)

	for _, ch := range lString {
		// проверка наличия текущего символа в мапе (если есть значит уже нет уникальности)
		if _, exists := cch[ch]; exists {
			return false
		}
		// добавляем чар в мапу
		cch[ch] = struct{}{}
	}

	return true
}

func main() {
	testLists := []string{"abcd", "abCdefAaf123", "aAbcd"}

	for _, lst := range testLists {
		res := isCharsUnique(lst)
		fmt.Printf("%s  %t\n", lst, res)
	}
}
