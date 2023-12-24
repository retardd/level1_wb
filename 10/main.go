package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для группировки температур
	groups := make(map[int][]float64)

	// Группируем температуры по шагу в 10 градусов
	for _, temp := range temperatures {
		groupKey := int(math.Floor(temp/10.0) * 10)
		groups[groupKey] = append(groups[groupKey], temp)
	}

	// Выводим результат
	for key, group := range groups {
		fmt.Printf("Группа для колебаний от %d д %d: %v\n", key, key+10, group)
	}
}
