package main

import (
	"fmt"
)

var justString []byte

// функция пытается присвоить срез v[:100] переменной justString.
// однако, если v имеет длину менее 100, мы выходим за его границы
// также если v будет крайне большой строкой, а justString останется в области видимости после выполнения
// someFunc() и будет ссылаться на массив данных, то он удержит его в памяти и
// не даст сборщику очистить место

func main() {
	v := createHugeString(1 << 10)
	if len(v) > 100 {
		justString = make([]byte, 100)
		copy(justString, v[:100])
	} else {
		justString = make([]byte, len(v))
		copy(justString, v)
	}

	fmt.Println(string(justString))
}

func createHugeString(size int) string {
	var b string
	i := 0
	for i < size {
		i += 14
		b += fmt.Sprintf("SymbolNumber%d", i)
	}
	return b
}
