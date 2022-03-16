package main

import (
	"full_version_pt/generator"
	"full_version_pt/printer"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

// const (
// 	defaultChannelValue = 2
// 	defaultFlagValue    = 10
// )

// Определяет размер буфера чтения и отправки соединение WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrader.CheckOrigin определит,
	//разрешено ли подключение входящего запроса из другого домена,
	//и если это не так, они будут поражены ошибкой CORS.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	//upgrader.Upgrade примет Response Writer и указатель на HTTP-запрос
	//и вернет нам указатель на соединение WebSocket или ошибку,
	//если не удалось выполнить Upgrade.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// ws.WriteMessage пишет ответ клиенту
	log.Println("Client Connected")

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

	genNum(flagValue, channelValue, ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/ws", wsEndpoint)
}

func main() {

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8000", nil))

	// // Парсинг флагов.
	// // maxValueFlag - Максимальное количество генерируемых чисел
	// maxValueFlag := flag.Int("max", defaultFlagValue, "the max value")
	// maxChannelFlag := flag.Int("chanel", defaultChannelValue, "the channel amount")

	// flag.Parse()
	// numberFlag := *maxValueFlag
	// channelFlag := *maxChannelFlag
	// // Проверка знака числа флага и на равенство 0.
	// if numberFlag < 1 || channelFlag < 0 {
	// 	// Если проверка не пройдена, то завершение программы. os.Exit(1)
	// 	fmt.Println("Было введено недопустимое значение. Попробуйте значение больше 0.")
	// 	os.Exit(1)
	// }

}

func genNum(numberFlag int, channelFlag int, ws *websocket.Conn) {
	var wg sync.WaitGroup

	// Инициализация каналов передачи рандомных чисел и готового среза
	numberChannel := make(chan int, numberFlag*10)

	// Инициализация принтера
	printer := printer.NewPrinter(numberFlag, ws)
	wg.Add(1)
	go printer.PrintNumbers(&wg, numberChannel)

	// Инициализация генераторов
	for i := 0; i < channelFlag; i++ {
		generator, generatorChan := generator.NewGenerator(numberFlag)
		printer.AppendGenerator(generatorChan)
		wg.Add(1)
		go generator.RandomNumber(&wg, numberChannel)
	}
	wg.Wait()
}
