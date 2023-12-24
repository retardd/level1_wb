package main

import "fmt"

// OldStruct старый интерфейс
type OldStruct struct{}

// LegacyMethod метод старого интерфейса
func (os *OldStruct) LegacyMethod() {
	fmt.Println("старый метод:(")
}

// NewInterface новый интерфейс
type NewInterface interface {
	NewMethod()
}

// NewStruct структура нового интерфейса
type NewStruct struct{}

// NewMethod новый метод
func (ns *NewStruct) NewMethod() {
	fmt.Println("новый метод:)")
}

// Adapter реализация адаптера для использования старого метода
type Adapter struct {
	OldStruct *OldStruct
}

// NewMethod реализация старого метода под видом нового, который реализует новый интерфейс
func (a *Adapter) NewMethod() {
	a.OldStruct.LegacyMethod()
	fmt.Println("...но под реализацией нового интерфейса")
}

func main() {
	// старая система
	oldStruct := &OldStruct{}
	oldStruct.LegacyMethod()

	// новая
	newStruct := &NewStruct{}
	newStruct.NewMethod()

	// Используем адаптер для старой системы, чтобы она работала с новым интерфейсом
	adapter := &Adapter{OldStruct: oldStruct}
	AdapterWorkCheck(adapter)
}

// AdapterWorkCheck функция принимает любой объект, реализующий интерфейс NewInterface
func AdapterWorkCheck(obj NewInterface) {
	obj.NewMethod()
}
