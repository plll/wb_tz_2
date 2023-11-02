package main

import (
	"fmt"
	"slices"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	stringSlice := []string{"cat", "cat", "dog", "cat", "tree"}
	newSet := createSetFromSlice(stringSlice)
	fmt.Println(newSet)
}

func createSetFromSlice(slice []string) []string {
	res := make([]string, 0)
	for _, x := range slice {
		if !slices.Contains(res, x) { // Последовательно идем по элементам слайса и добавляем все элементы, которые не находим в res
			res = append(res, x)
		}
	}
	return res
}
