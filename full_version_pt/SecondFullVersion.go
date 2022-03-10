package main

import (
	"flag"
	"fmt"
	"full_version_pt/generator"
	"full_version_pt/printer"
	"os"
	"sync"
)

// defaultFlagValue - Дефолтное количество генерируемых чисел
const (
	defaultFlagValue   = 11
	defaultChanelValue = 2
)

func main() {
	// Парсинг флагов.
	// maxValueFlag - Максимальное количество генерируемых чисел
	maxValueFlag := flag.Int("max", defaultFlagValue, "the max value")
	numberFlag := *maxValueFlag
	maxChanelFlag := flag.Int("chanel", defaultChanelValue, "the chanel amount")
	chanelFlag := *maxChanelFlag
	flag.Parse()

	// Проверка знака числа флага и на равенство 0.
	if numberFlag < 1 || chanelFlag < 0 {
		// Если проверка не пройдена, то завершение программы. os.Exit(1)
		fmt.Println("Было введено недопустимое значение. Попробуйте значение больше 0.")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	// Инициализация каналов передачи рандомных чисел и готового среза
	numberChanel := make(chan int, numberFlag)
	resultChanel := make(chan []int, 1)

	// Инициализация принтера
	printer := printer.NewPrinter(numberFlag)
	wg.Add(1)
	go printer.PrintNumbers(&wg, numberChanel, resultChanel)

	// Инициализация генераторов
	for i := 0; i < chanelFlag; i++ {
		generator, generatorChan := generator.NewGenerator(numberFlag)
		printer.AppendGenerator(generatorChan)
		wg.Add(1)
		go generator.RandomNumber(&wg, numberChanel)
	}

	resultNumbers := <-resultChanel
	wg.Wait()
	fmt.Println(resultNumbers)
}
