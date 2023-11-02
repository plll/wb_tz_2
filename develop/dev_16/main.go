package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func partition(arr []int, low, high int) ([]int, int) { // Выбор опорного элемента и сдвиг всех меньших элементов левее опорного
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int { // Обертка над partion, которая является функцией для рекурсивного вызова самой себя и алгоритма сортировки
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func main() {
	s := []int{1, 2, 5, 3, 2, 7, 4, 8}
	sSorted := quickSort(s, 0, len(s)-1)
	fmt.Println(sSorted)
}
