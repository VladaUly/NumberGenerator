package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Создание каналов
	var c chan [100]int = make(chan [100]int)
	var d chan []int = make(chan []int)

	// Горутины
	go numberGenerator(c)
	go removeDuplicateElement(c, d)
	go numberOutput(d)
}

// Объект, поддерживающий поток генерации чисел
func numberGenerator(c chan [100]int) {
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

	c <- s
}

// Объект, поддерживающий поток вывода чисел
func numberOutput(d chan []int) {
	// Вывод дедублицированных чисел в консоль
	nmbrs := <-d
	fmt.Println(nmbrs)
}

// Удаление дубликатов срезов массива
func removeDuplicateElement(c chan [100]int, d chan []int) {
	array := <-c
	// Создание среза
	result := make([]int, 0, len(array))
	// Создание карты
	temp := map[int]struct{}{}
	for _, item := range array {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			//Создание нового среза из уже существующих
			result = append(result, item)
		}
	}
	d <- result

}
