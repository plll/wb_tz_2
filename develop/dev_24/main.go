package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func findDistance(p1, p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)) // Формула нахождения расстояния между двумя точками
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(3, 3)
	dist := findDistance(p1, p2)
	fmt.Println(dist)
}
