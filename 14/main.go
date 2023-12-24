package main

import (
	"fmt"
)

func main() {

	var someValue interface{} = 42

	// Целочисленный
	printType(someValue)

	// Строковый
	someValue = "hello"
	printType(someValue)

	// Логический
	someValue = true
	printType(someValue)

	// Канал
	someValue = make(chan int)
	printType(someValue)
}

func printType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan int:
		fmt.Println("Тип: channel")
	default:
		fmt.Println("Тип: неизвестный")
	}
}
