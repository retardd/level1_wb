package main

import (
	"fmt"
	"sync"
)

func main() {
	// Заданный массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Создание WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Итерация по массиву чисел
	for _, num := range numbers {
		// Увеличение счетчика ожидаемых горутин
		wg.Add(1)

		// Запуск горутины для вычисления квадрата числа
		go func(n int) {
			defer wg.Done()
			res := n * n
			// Вывод квадрата в stdout
			fmt.Printf("квадрат числа %d: %d\n", n, res)
		}(num)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
}
