package main

import "fmt"

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	ch := make(chan int)
	mass := []int{2, 4, 6, 8, 10}
	for _, x := range mass {
		go getSquare(x, ch) //Запускаем по горутине для просчета квадрата числа
	}
	squareSum := 0
	for i := 0; i < 5; i++ {
		squareSum += <-ch //В цикле получаем квадраты и добавляем их к переменной суммы
	}
	fmt.Println(squareSum)
}

func getSquare(x int, ch chan int) {
	ch <- x * x //Отправляем квадраты чисел через канал
}
