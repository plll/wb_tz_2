package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Println(checkIfSymbolsUnique(s1))
	fmt.Println(checkIfSymbolsUnique(s2))
	fmt.Println(checkIfSymbolsUnique(s3))
}

func checkIfSymbolsUnique(s string) bool {
	runeSlice := []rune(strings.ToLower(s)) // Приводим к нижнему регистру и разбиваем строку на слайс рун
	for i := 0; i < len(runeSlice); i++ {
		for j := 0; j < len(runeSlice); j++ {
			if i != j && runeSlice[i] == runeSlice[j] {
				return false
			}
		}
	}
	return true
}
