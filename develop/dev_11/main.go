package main

import (
	"fmt"
	"slices"
)

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	set1 := []int{1, 3, 4}
	set2 := []int{1, 2, 3, 6}
	fmt.Println(findCommonInSets(set1, set2))
}

func findCommonInSets[T comparable](set1, set2 []T) []T {
	res := make([]T, 0)
	if len(set1) > len(set2) {
		for _, x := range set1 {
			if slices.Contains(set2, x) { // Проверяем, что элемент первого слайса содержится во втором, можно было бы также проверять нет ли в res уже этого элемента, но нам дано что рассматриваются множества
				res = append(res, x)
			}
		}
	} else {
		for _, x := range set2 {
			if slices.Contains(set1, x) {
				res = append(res, x)
			}
		}
	}
	return res
}
