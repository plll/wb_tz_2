package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	s := "snow dog sun"
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	words := strings.Fields(s)          // Получаем слайс слов
	for i := 0; i < len(words)/2; i++ { // Разворачиваем слайс слов
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ") // Соединяем слова из слайса через пробел
}
