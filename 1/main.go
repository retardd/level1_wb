package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) SayHello() {
	fmt.Printf("меня зовут %s и мне %d лет.\n", h.name, h.age)
}

type Action struct {
	Human // Встраивание структуры Human
}

// SaySomething Новый метод в структуре Action
func (a *Action) SaySomething() {
	fmt.Printf("я - %s\n", a.name)
}

func main() {
	person := Action{
		Human: Human{
			name: "Иван",
			age:  22,
		},
	}

	// Вызов методов как из Human, так и из Action
	person.SayHello()
	person.SaySomething()
}
