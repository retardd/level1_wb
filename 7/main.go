package main

import (
	"fmt"
	"sync"
)

// структура с мютексом и мапой для безопасной конкурентной записи
type mapMutex struct {
	sync.Mutex
	m map[int]string
}

func main() {
	//создаем все переменные
	var wg sync.WaitGroup
	SafeMap := mapMutex{m: make(map[int]string)}

	// конкурентная запись в мапу через 5 горутин
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(SafeMap *mapMutex, key int, value string) {
			SafeMap.Lock()
			//после завершения горутины структура откроется
			defer SafeMap.Unlock()
			defer wg.Done()

			SafeMap.m[key] = value

			fmt.Printf("[%d] = %s\n", key, value)
		}(&SafeMap, i, fmt.Sprintf("some_string_n%d", i))
	}

	wg.Wait()
	// вывод мапа
	fmt.Println("Final Map:", SafeMap.m)
}
