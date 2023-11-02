package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	sRune := []rune(s) // Кастим стринг в слайс рун, чтобы итерироваться по нему
	for i := 0; i < len(sRune)/2; i++ {
		sRune[i], sRune[len(sRune)-i-1] = sRune[len(sRune)-i-1], sRune[i]
	} // В цикле меняем местами последний с первым и тд по возврастанию i
	return string(sRune)
}
