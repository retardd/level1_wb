package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем канал для передачи данных
	dataChan := make(chan int)

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Запускаем горутину для записи данных в канал
	wg.Add(1)
	go sendData(dataChan, &wg)

	// Запускаем горутину для чтения данных из канала
	wg.Add(1)
	go receiveData(dataChan, &wg)

	// Ждем N секунд
	N := 5
	time.Sleep(time.Duration(N) * time.Second)

	// Закрываем канал для завершения горутин
	close(dataChan)

	// Ожидаем завершения горутин
	wg.Wait()

	fmt.Println("Программа завершена.")
}

// Функция для записи данных в канал
func sendData(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		// Отправляем значение в канал
		ch <- i
		// Задержка перед следующей отправкой
		time.Sleep(1 * time.Second)
	}
}

// Функция для чтения данных из канала
func receiveData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Читаем значение из канала
		data, ok := <-ch
		if !ok {
			// Канал закрыт, завершаем чтение
			fmt.Println("Чтение данных завершено.")
			return
		}
		fmt.Printf("Принято значение: %d\n", data)
	}
}
