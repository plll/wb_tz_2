package main

import (
	"fmt"
	"log"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func changeIBit(x int64, i, newBit int) int64 {
	if newBit == 1 {
		newBit = 49
	} else {
		newBit = 48
	}
	binSlice := []rune(fmt.Sprintf("%b", x))                 // Кастим инт в двочиную сс и разбиваем строку на слайс
	binSlice[i] = rune(newBit)                               // заменяем нужный индекс на новый бит
	result, err := strconv.ParseInt(string(binSlice), 2, 64) // Конвертим строку битов в инт
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func main() {
	var x int64 = 34567890123123123
	i := 6
	fmt.Println(x)
	x = changeIBit(x, i, 1)
	fmt.Println(x)
	x = changeIBit(x, i, 0)
	fmt.Println(x)
}
