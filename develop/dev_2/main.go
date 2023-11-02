package main

import "fmt"

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	ch := make(chan int)
	mass := []int{2, 4, 6, 8, 10}
	for _, x := range mass {
		go getSquare(x, ch) //Запускаем по горутине для просчета квадрата числа
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch) //В цикле получаем квадраты и выводим
	}
}

func getSquare(x int, ch chan int) {
	ch <- x * x //Отправляем квадраты чисел через канал
}
