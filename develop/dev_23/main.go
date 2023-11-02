package main

import (
	"fmt"
)

// Удалить i-ый элемент из слайса.

func removeIndex[T any](s []T, index int) []T { // Создаем функцию готовую работать с слайсом, внутри которого будет любой тип данных
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("До удаления: ", slice) //[0 1 2 3 4 5 6 7 8 9]
	slice = removeIndex(slice, 5)
	fmt.Println("После удаления: ", slice)
}
