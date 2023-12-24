package main

import (
	"fmt"
)

func main() {
	// Создаем каналы для передачи данных между горутинами
	input := make(chan int)
	stdout := make(chan int)

	// Запускаем горутину для записи чисел в канал inputChan
	go func() {
		numbers := []int{1, 2, 3, 4, 5}

		for _, num := range numbers {
			input <- num
		}

		close(input)
	}()

	// Запускаем горутину для умножения чисел на 2 и передачи результата в канал outputChan
	go func() {
		for num := range input {
			result := num * 2
			stdout <- result
		}

		close(stdout)
	}()

	// Выводим результат из канала outputChan в stdout
	func() {
		for result := range stdout {
			fmt.Println(result)
		}
	}()
}
