package main

import (
	"flag"
	"fmt"
	"os"
	""
)

// defaultFlagValue - Дефолтное количество генерируемых чисел
const (
	defaultFlagValue   = 100
	defaultChanelValue = 2
)

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

}
