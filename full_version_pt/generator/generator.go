package generator

import (
	"math/rand"
	"sync"
	"time"
)

type Generator struct {
	limit int
}

func NewGenerator(limit int) *Generator {
	return &Generator{limit: limit}
}

func (g *Generator) RandomNumber(wg sync.WaitGroup, number chan<- int, readySignal chan int) {

	rand.Seed(time.Now().UnixNano())
	for {
		ready := <-readySignal
		if ready == 1 {
			wg.Done()
		}
		randomNumber := rand.Intn(g.limit + 1)
		number <- randomNumber

	}

}
