package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2

	if index >= len(slice) {
		fmt.Println("out of range")
	} else {
		// удаление элемента по индексу (реализовал с сохранением порядка слайса)
		slice := append(slice[:index], slice[index+1:]...)
		fmt.Println(slice)
	}

}
