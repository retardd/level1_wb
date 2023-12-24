package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Создаем канал для передачи данных от главной горутины к воркерам
	dataChan := make(chan int)

	// Количество воркеров вводится из консоли
	var numWorkers int
	fmt.Print("количество воркеров: ")
	_, err := fmt.Scan(&numWorkers)
	if err != nil {
		return
	}

	// Генератор случайных чисел
	rand.Seed(100)

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Инициализация воркеров (с 1 по numWorkers)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int, dataChan <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for data := range dataChan {
				fmt.Printf("воркер %d: %d\n", id, data)
			}
			fmt.Printf("вокрек %d отключен.\n", id)
		}(i, dataChan, &wg)
	}

	// Запускаем горутину для записи данных в канал
	go func(dataChan chan<- int) {
		for {
			// Генерируем случайное число и в канал
			data := rand.Intn(101)
			dataChan <- data
			// Задержка перед следующей итерацией
			time.Sleep(time.Second)
		}
	}(dataChan)

	// Канал для отслеживания сигнала завершения программы
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Ожидаем получения сигнала завершения
	<-signalChan

	// Закрываем канал данных
	close(dataChan)

	// Ожидаем завершения всех воркеров
	wg.Wait()
}
