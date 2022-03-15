package main

import (
	"fmt"
	"full_version_pt/generator"
	"full_version_pt/printer"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/generator", func(w http.ResponseWriter, r *http.Request) {
		numberURLFlag := r.URL.Query().Get("number")
		channelURLFlag := r.URL.Query().Get("chanel")

		flagValue, err := strconv.Atoi(numberURLFlag)
		if err != nil {
			panic(err)
		}
		channelValue, err := strconv.Atoi(channelURLFlag)
		if err != nil {
			panic(err)
		}

		numberFlag := flagValue
		channelFlag := channelValue

		result := generatorFunc(numberFlag, channelFlag)

		for i := 0; i < len(result); i++ {
			intResult := result[i]
			byteResult := []byte(strconv.FormatInt(int64(intResult), 10))
			w.Write(byteResult)
			space := []byte(" ")
			w.Write(space)
		}
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
func generatorFunc(numberFlag int, channelFlag int) []int {
	// // Парсинг флагов.
	// // maxValueFlag - Максимальное количество генерируемых чисел
	// maxValueFlag := flag.Int("max", defaultFlagValue, "the max value")
	// maxChannelFlag := flag.Int("chanel", defaultChannelValue, "the channel amount")

	// flag.Parse()
	// numberFlag := *maxValueFlag
	// channelFlag := *maxChannelFlag
	// fmt.Println(*maxValueFlag)
	// // Проверка знака числа флага и на равенство 0.
	// if numberFlag < 1 || channelFlag < 0 {
	// 	// Если проверка не пройдена, то завершение программы. os.Exit(1)
	// 	fmt.Println("Было введено недопустимое значение. Попробуйте значение больше 0.")
	// 	os.Exit(1)
	// }

	var wg sync.WaitGroup

	// Инициализация каналов передачи рандомных чисел и готового среза
	numberChannel := make(chan int, numberFlag*10)
	resultChannel := make(chan []int, 1)

	// Инициализация принтера
	printer := printer.NewPrinter(numberFlag)
	wg.Add(1)
	go printer.PrintNumbers(&wg, numberChannel, resultChannel)

	// Инициализация генераторов
	for i := 0; i < channelFlag; i++ {
		generator, generatorChan := generator.NewGenerator(numberFlag)
		printer.AppendGenerator(generatorChan)
		wg.Add(1)
		go generator.RandomNumber(&wg, numberChannel)
	}

	resultNumbers := <-resultChannel
	wg.Wait()
	fmt.Println(resultNumbers)
	return resultNumbers
}
