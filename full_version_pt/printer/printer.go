package printer

import (
	"sort"
	"sync"
)

// Printer записывает новые экземпляры принтера
type Printer struct {
	limit      int
	generators []chan string
}

// NewPrinter возвращает ссылку на новый экземпляр принтера случайных чисел
func NewPrinter(limit int) *Printer {
	return &Printer{limit: limit}
}

// AppendGenerator добавляет канал комманд генератора в срез generators
func (p *Printer) AppendGenerator(ch chan string) {
	p.generators = append(p.generators, ch)
}

// PrintNumbers запускает основной рабочий цикл принтера
func (p *Printer) PrintNumbers(wg *sync.WaitGroup, number <-chan int, resultChanel chan<- []int) {

	resultSlice := make([]int, 0)
	tempMap := map[int]struct{}{}
	for len(resultSlice) != p.limit {
		randomNumber := <-number
		if randomNumber == 0 {
			continue
		}
		if _, ok := tempMap[randomNumber]; !ok {
			tempMap[randomNumber] = struct{}{}
			resultSlice = append(resultSlice, randomNumber)
		}

	}
	sort.Ints(resultSlice)
	// Передача команды "СТОП" ВСЕМ существующим генераторам
	for _, generatorChan := range p.generators {
		generatorChan <- "stop"
	}
	resultChanel <- resultSlice
	wg.Done()
}
