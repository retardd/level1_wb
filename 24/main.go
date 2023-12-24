package main

import (
	"fmt"
	"math"
)

// Point структура с координатами точки
type Point struct {
	x, y float64
}

// NewPoint конструктор
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance вычисляет расстояние между двумя точками
func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	// корень квадратный из суммы квадратов разностей x и y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// две точки
	p1 := NewPoint(1, 6)
	p2 := NewPoint(4, 2)

	d := Distance(p1, p2)

	// stdout (2 знака после запятой)
	fmt.Printf("расстояни: %.2f\n", d)
}
