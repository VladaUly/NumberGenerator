package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
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

	// initialArray - Массив для преобразования в срез и добавления рандомных чисел
	var initialArray []int

	// Генерация ста случайных чисел от 0 до max
	for i := 0; i <= *maxValueFlag; i++ {
		randomNumber := rand.Intn(*maxValueFlag)

		if randomNumber != 0 {
			// Добавление рандомных чисел в массив
			initialArray = append(initialArray, randomNumber)
		}
	}

	// sortedNumbers - Переменная с отсортированным срезом
	sortedNumbers := removeDuplicateElement(initialArray)
	sort.Ints(sortedNumbers)

	// Вывод сгенерированных чисел в терминал
	fmt.Println(sortedNumbers)
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
