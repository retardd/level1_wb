package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	val int
	sync.Mutex
}

// Increment инкрементирование счетчика
func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.val++
}

// GetValue вывод текущего значение в stdout
func (c *Counter) GetValue() {
	c.Lock()
	defer c.Unlock()
	fmt.Println("счетчик:", c.val)
}

func main() {
	var wg sync.WaitGroup

	counter := Counter{}

	// 10 горутин для инкрементации счетчика
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// 10 раз инкрементируем значение в одной горутине (+10)
			for j := 0; j < 10; j++ {
				counter.Increment()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// Выводим итоговое значение счетчика
	counter.GetValue()
}
