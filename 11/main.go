package main

import (
	"fmt"
)

func main() {
	// Неупорядоченные множества
	set1 := []int{1, 2, 3, 4}
	set2 := []int{3, 4, 5, 6}

	intersectionSet := intersection(set1, set2)

	fmt.Println("Пересечение множеств:", intersectionSet)
}

// intersection находит пересечение двух множеств
func intersection(set1, set2 []int) []int {
	var result []int

	// Проходим по элементам первого множества
	for _, element := range set1 {
		// Проверяем, присутствует ли элемент во втором множестве
		for _, value := range set2 {
			if value == element {
				result = append(result, element)
			}
		}
	}

	return result
}
