package main

import (
	"fmt"
)

func main() {
	// 42 - 101010
	var number int64 = 42

	// 4-й бит в 1
	setBitVar := setBit(number, 4, 1)
	fmt.Printf("Число после установки 3-го бита в 1: %d\n", setBitVar)

	//2-й бит в 0
	setBitVar = setBit(number, 2, 0)
	fmt.Printf("Число после установки 2-го бита в 0: %d\n", setBitVar)
}

// устанавливает i-й бит в value (1 или 0) и возвращает результат.
func setBit(number int64, position uint, value int) int64 {
	// очищаем i-й бит, используя маску с битом, равным 1 на позиции i (все ост. позиции равны 0).
	clearMask := ^(int64(1) << position)
	number &= clearMask

	// i-й бит в нужное значение.
	number |= int64(value) << position

	return number
}
