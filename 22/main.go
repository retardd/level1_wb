package main

import (
	"fmt"
	"math/big"
)

func main() {
	// math/big
	a := new(big.Int)
	b := new(big.Int)

	// устанавливаем значения чисел (строка с числом, его с/c)
	a.SetString("33554432", 10)
	b.SetString("10000000", 10)

	// умножение
	result1 := new(big.Int)
	result1.Mul(a, b)
	fmt.Printf("умножение: %s\n", result1.String())

	// деление (соотв. без остатка)
	result2 := new(big.Int)
	result2.Div(a, b)
	fmt.Printf("деление: %s\n", result2.String())

	// сложение
	result3 := new(big.Int)
	result3.Add(a, b)
	fmt.Printf("сложение: %s\n", result3.String())

	// вычитание
	result4 := new(big.Int)
	result4.Sub(a, b)
	fmt.Printf("вычитание: %s\n", result4.String())

}
