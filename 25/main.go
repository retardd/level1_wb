package main

import (
	"fmt"
	"time"
)

// mySleep - если я праивльно понял суть реализации
func mySleep(s int) {
	dur := time.Duration(s) * time.Second
	time.Sleep(dur)
}

func main() {
	s := 3
	fmt.Printf("поспим %d секунд\n", s)
	mySleep(s)
	fmt.Println("поспали:)")
}
