package main

import (
	"fmt"
	"sort"
)

// для других типов данных нужно переопеределять интерфейс для библиы sort

func main() {
	arr := []int{5, 3, 8, 2, 1, 7, 4, 6}

	// sort.IntSlice для преобразования слайса в data Interface
	sort.Sort(sort.IntSlice(arr))
	//.......(sort.Reverse(sort.IntSlice(arr))

	// Вывод отсортированного массива
	fmt.Println(arr)
}
