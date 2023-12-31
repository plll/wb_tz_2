package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := gen(x)
	out := sq(c)

	// Выводим значения.
	for i := 0; i < len(x); i++ {
		fmt.Println(<-out)
	}
}

func gen(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
