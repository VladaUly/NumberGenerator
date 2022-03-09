package main

import (
	"flag"
	"fmt"
	"full_version_pt/generator"
	"full_version_pt/printer"
	"os"
)

// defaultFlagValue - Дефолтное количество генерируемых чисел
const (
	defaultFlagValue   = 10
	defaultChanelValue = 2
)

func main() {
	fmt.Println("Ввведите max количество генерируемых чисел и каналов ( chanel )")
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

	numberChanel := make(chan int)
	resultChanel := make(chan []int, 1)

	printer := printer.NewPrinter(numberFlag)
	go printer.PrintNumbers(numberChanel, resultChanel)

	for i := 0; i < chanelFlag; i++ {
		generator := generator.NewGenerator(numberFlag)
		go generator.RandomNumber(numberChanel)
	}
	resultNumbers := <-resultChanel
	fmt.Println(resultNumbers)
}
