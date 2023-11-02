package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	arr := []int{0, 2, 5, 6, 8, 11, 14}
	index, ok := binarySearch(arr, 14)
	if ok {
		fmt.Println(index)
	} else {
		fmt.Println("Не найдено")
	}
}

func binarySearch(arr []int, key int) (int, bool) {
	low := 0
	high := len(arr) - 1

	for low <= high { // На каждой итерации проверяем середину диапозона с числом, которое ищем, если число меньше - то делаем диапозон [low, mid-1], иначе [mid+1, high]
		mid := (low + high) / 2

		if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(arr) || arr[low] != key {
		return -1, false
	}

	return low, true
}
