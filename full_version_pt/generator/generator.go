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

func (g *Generator) RandomNumber(number chan<- int, mutex chan sync.Mutex) {
	mu := <-mutex
	mu.Lock()
	defer mu.Lock()
	rand.Seed(time.Now().UnixNano())
	for {
		randomNumber := rand.Intn(g.limit + 1)
		number <- randomNumber

	}
}
