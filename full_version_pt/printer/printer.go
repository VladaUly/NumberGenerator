package printer

import (
	"sort"
	"sync"
)

type Printer struct {
	limit int
}

func NewPrinter(limit int) *Printer {
	return &Printer{limit: limit}
}

func (p *Printer) PrintNumbers(wg sync.WaitGroup, number <-chan int, resultChanel chan<- []int, readySignal chan int) {

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
	ready := 1
	readySignal <- ready
	sort.Ints(resultSlice)
	resultChanel <- resultSlice
	wg.Done()
}
