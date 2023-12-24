package main

import "fmt"

func main() {
	int1 := 340123
	int2 := -23102
	fmt.Printf("int1: %d; int2: %d\n", int1, int2)
	int1 = int2 - int1
	int2 = int2 - int1
	int1 = int1 + int2
	fmt.Printf("int1: %d; int2: %d\n", int1, int2)
	return
}
