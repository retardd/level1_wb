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

	resultChannel := make(chan int, 5)

	// Итерация по массиву чисел
	for _, num := range numbers {
		// Увеличение счетчика ожидаемых горутин
		wg.Add(1)

		// Запуск горутины для вычисления квадрата числа
		go func(n int) {
			defer wg.Done()
			res := n * n
			// Вывод квадрата в канал resultChannel
			resultChannel <- res
		}(num)
	}

	wg.Wait()
	// Закрытие канала после завершения всех горутин
	close(resultChannel)

	var sum int
	for res := range resultChannel {
		sum += res
	}

	// Вывод итоговой суммы квадратов
	fmt.Printf("сумма квадратов: %d\n", sum)
}
