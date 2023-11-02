package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	var a = 5
	var b = 10
	fmt.Println("До обмена: a =", a, ", b =", b) // Алгоритм обмена при помощи исключающего ИЛИ
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("После обмена: a =", a, ", b =", b)
}
