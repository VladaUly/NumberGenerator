package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// defaultFlagValue - Дефолтное количество генерируемых чисел
const (
	defaultFlagValue = 100
)

func main() {
	// maxValueFlag - Максимальное количество генерируемых чисел
	maxValueFlag := flag.Int("max", defaultFlagValue, "the max value")
	// Парсинг
	flag.Parse()

	// initialArray - Массив для преобразования в срез
	var initialArray []int

	// Генерация ста случайных чисел от 0 до max
	for i := 0; i <= *maxValueFlag-1; i++ {
		randomNumber := rand.Intn(*maxValueFlag)

		// Добавление рандомных чисел в массив
		initialArray = append(initialArray, randomNumber)
	}

	// Вывод сгенерированных чисел в терминал
	fmt.Println(removeDuplicateElement(initialArray))
}

// removeDuplicateElement - Удаление дубликатов срезов массива
func removeDuplicateElement(sliceNumbers []int) []int {
	//resultSlice - Создание среза из массива initialArray
	resultSlice := make([]int, 0, len(sliceNumbers))
	//rempMap - Создание карты
	tempMap := map[int]struct{}{}
	for _, item := range sliceNumbers {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = struct{}{}
			resultSlice = append(resultSlice, item)
		}
	}
	return resultSlice
}
