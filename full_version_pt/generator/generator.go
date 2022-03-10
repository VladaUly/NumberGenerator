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

func (g *Generator) RandomNumber(number chan<- int, wg chan sync.WaitGroup) {
	waitGroup := <-wg
	rand.Seed(time.Now().UnixNano())
	for {
		randomNumber := rand.Intn(g.limit + 1)
		number <- randomNumber
		waitGroup.Done()
		wg <- waitGroup
	}

}
