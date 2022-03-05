package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

// defaultFlagValue - Дефолтное количество генерируемых чисел
const (
	defaultFlagValue = 100
)

func main() {

	// 1.Парсинг флагов.

	// maxValueFlag - Максимальное количество генерируемых чисел
	maxValueFlag := flag.Int("max", defaultFlagValue, "the max value")
	flag.Parse()

	// 2.Проверка знака числа флага и на равенство 0.
	if *maxValueFlag < 1 {
		// 3.Если проверка не пройдена, то завершение программы. os.Exit(1)
		fmt.Println("Было введено недопустимое значение. Попробуйте значение больше 0.")
		os.Exit(1)
	}

	// resultSlice - Создание среза из массива initialArray
	resultSlice := make([]int, 0)
	// tempMap - Создание ассоциативного массива
	tempMap := map[int]struct{}{}
	// 4.Начало цикла.
	// 5.Если длина среза с числами не равна флагу диапазона, то
	for *maxValueFlag != len(resultSlice) {
		// 6.Сгенерировать случайное число.
		randomNumber := rand.Intn(*maxValueFlag + 1)
		if randomNumber == 0 {
			continue
		}
		// 7.Проверить, не было ли числе УЖЕ сгенерировано.
		if _, ok := tempMap[randomNumber]; !ok {
			tempMap[randomNumber] = struct{}{}
			// 9.Если нет, то добавить сгенерированное число в срез чисел.
			resultSlice = append(resultSlice, randomNumber)

			// 8.Если было. то continue.
		} else {
			continue
		}

		// 10.Конец цикла
	}
	// 11.Сортировка среза
	sort.Ints(resultSlice)
	// 12.Вывод среза в терминал.
	fmt.Println(resultSlice)
	// 13.Завершение программы.  os.Exit(0)
	os.Exit(0)

}
