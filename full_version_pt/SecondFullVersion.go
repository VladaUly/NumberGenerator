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
	defaultFlagValue   = 10
	defaultChanelValue = 2
)

var wg sync.WaitGroup

func main() {
	// 1.Парсинг флагов.
	// maxValueFlag - Максимальное количество генерируемых чисел
	maxValueFlag := flag.Int("max", defaultFlagValue, "the max value")
	numberFlag := *maxValueFlag
	maxChanelFlag := flag.Int("chanel", defaultChanelValue, "the chanel amount")
	chanelFlag := *maxChanelFlag
	flag.Parse()

	// 2.Проверка знака числа флага и на равенство 0.
	if numberFlag < 1 || chanelFlag < 0 {
		// 3.Если проверка не пройдена, то завершение программы. os.Exit(1)
		fmt.Println("Было введено недопустимое значение. Попробуйте значение больше 0.")
		os.Exit(1)
	}

	numberChanel := make(chan int, 1)
	resultChanel := make(chan []int, 1)
	readySignal := make(chan int)

	printer := printer.NewPrinter(numberFlag)
	wg.Add(1)
	go printer.PrintNumbers(wg, numberChanel, resultChanel, readySignal)

	for i := 0; i < chanelFlag; i++ {
		generator := generator.NewGenerator(numberFlag)
		wg.Add(1)
		go generator.RandomNumber(wg, numberChanel, readySignal)
	}

	resultNumbers := <-resultChanel
	wg.Wait()
	fmt.Println(resultNumbers)
}
