package main

import (
	"fmt"
	"sort"
)

func main() {
	// пример сортированного среза
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// искомый элемент
	target := 2

	// выполняем бинарный поиск
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	// вывод (index < len(arr) чтобы избежать случайного падения от (index out of range))
	if index < len(arr) && arr[index] == target {
		fmt.Printf("элемент %d найден под индексом %d\n", target, index)
	} else {
		fmt.Printf("элемент %d не найден\n", target)
	}
}
