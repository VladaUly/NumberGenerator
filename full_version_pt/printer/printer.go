package printer

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

// Printer записывает новые экземпляры принтера
type Printer struct {
	limit      int
	generators []chan string
	webSocket  *websocket.Conn
}

// NewPrinter возвращает ссылку на новый экземпляр принтера случайных чисел
func NewPrinter(limit int, ws *websocket.Conn) *Printer {
	return &Printer{limit: limit, webSocket: ws}
}

// AppendGenerator добавляет канал комманд генератора в срез generators
func (p *Printer) AppendGenerator(ch chan string) {
	p.generators = append(p.generators, ch)

}

// PrintNumbers запускает основной рабочий цикл принтера
func (p *Printer) PrintNumbers(wg *sync.WaitGroup, number chan int) {
	resultSlice := make([]int, 0)
	tempMap := map[int]struct{}{}
	neededNum := 1
	for {
		randomNumber := <-number
		if len(resultSlice) == p.limit {
			number = nil
			// Передача команды "СТОП" ВСЕМ существующим генераторам
			for _, generatorChan := range p.generators {
				generatorChan <- "stop"

			}
			break
		}

		if randomNumber == 0 {
			continue
		}
		if _, ok := tempMap[randomNumber]; !ok {
			if randomNumber == neededNum {
				tempMap[randomNumber] = struct{}{}
				resultSlice = append(resultSlice, randomNumber)
				neededNum++

				byteResult := []byte(strconv.FormatInt(int64(randomNumber), 10))
				p.webSocket.WriteMessage(1, []byte(byteResult))
				fmt.Println(randomNumber)

			}
		}

	}

	wg.Done()
}
