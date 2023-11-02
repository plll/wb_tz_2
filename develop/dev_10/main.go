package main

import (
	"fmt"
	"math"
	"slices"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//
//
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
//(Пример данный в задаче неправильный, тк -20 [-30;-20], -10 [-20, -10], 0 [-10;0], 10 [0; 10] Исходя из условия должно быть так, но в примере неправильно рассчитан первый диапозон

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5, 0.5}
	res := splitTemperatures(temp)
	keys := make([]float64, 0, len(res))
	for k := range res {
		keys = append(keys, k)
	}
	// Сортируем ключи мапы
	slices.Sort(keys)

	// Выводим диапозоны в отсортированном виде
	for _, k := range keys {
		fmt.Printf("%v: %v, ", k, res[k])
	}
}

func splitTemperatures(temp []float64) map[float64][]float64 {
	minimum := slices.Min(temp)
	maximum := slices.Max(temp)
	low := math.Round(minimum/10) * 10        // Находим нижнюю границу
	high := (math.Round(maximum/10) + 1) * 10 // Находим верхнюю границу
	res := make(map[float64][]float64)
	for i := low; i <= high; i += 10 { // Итерируемся по границам от минимума до максимуа с шагом 10
		line := make([]float64, 0)       // Создаем слайс в который будем добавлять числа входящие в диапозон
		for j := 0; j < len(temp); j++ { // Проверяем все числа на входимость в диапозон в цикле
			if i <= temp[j] && temp[j] < i+10 {
				line = append(line, temp[j])
			}
		}
		if len(line) != 0 {
			res[i] = line
		}
	}
	return res
}
