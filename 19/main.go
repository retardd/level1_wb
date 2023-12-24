package main

import "fmt"

func main() {
	SomeString := "1. привет - тевирп .2"
	var ReverseString string
	runes := []rune(SomeString)
	tempLen := len(runes)

	for tempLen > 0 {
		tempLen--
		tempChar := runes[tempLen : tempLen+1]
		ReverseString = fmt.Sprintf("%s%s", ReverseString, string(tempChar))

	}

	fmt.Println(ReverseString)
	return
}
