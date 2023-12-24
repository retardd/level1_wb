package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// закрытие канала для завершения горутины
	stopChan := make(chan int)

	wg.Add(1)
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("go exit")
				wg.Done()
				return
			default:
				fmt.Println("go work")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)

	// закрываем канал
	close(stopChan)
	wg.Wait()

	// использование контекста
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("go exit")
				wg.Done()
				return
			default:
				fmt.Println("go work")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)

	// тмена контекста
	cancel()
	wg.Wait()
}
