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

func (p *Printer) PrintNumbers(number <-chan int, resultChanel chan<- []int, wg chan sync.WaitGroup) {
	waitGroup := <-wg
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
	resultChanel <- resultSlice
	waitGroup.Done()
	wg <- waitGroup

}
