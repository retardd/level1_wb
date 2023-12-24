package main

import "fmt"

func main() {
	listStrings := []string{"cat", "cat", "dog", "cat", "tree"}

	setStrings := createSet(listStrings)

	fmt.Println("Множество:", setStrings)
}

func createSet(sequence []string) []string {
	set := make([]string, 0)
	seen := make(map[string]bool)

	for _, element := range sequence {
		// Если элемент не встречался ранее, добавляем его в множество и отмечаем, что он был виден
		if !seen[element] {
			set = append(set, element)
			seen[element] = true
		}
	}

	return set
}
