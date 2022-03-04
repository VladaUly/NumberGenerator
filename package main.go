package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Определение флагов
	maxp := flag.Int("max", 100, "the max value")
	// Парсинг
	flag.Parse()

	var s [100]int

	// Генерация ста случайных чисел от 0 до max
	for i := 0; i <= 99; i++ {
		x := rand.Intn(*maxp)

		s[i] = x
	}

	//
	fmt.Println(removeDuplicateElement(s))
}

// Удаление дубликатов срезов массива
func removeDuplicateElement(number [100]int) []int {
	result := make([]int, 0, len(number))
	temp := map[int]struct{}{}
	for _, item := range number {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
