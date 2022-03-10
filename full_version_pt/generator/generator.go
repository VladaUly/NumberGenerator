package generator

import (
	"math/rand"
	"time"
)

type Generator struct {
	limit int
}

func NewGenerator(limit int) *Generator {
	return &Generator{limit: limit}
}

func (g *Generator) RandomNumber(number chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for {
		randomNumber := rand.Intn(g.limit + 1)
		number <- randomNumber

	}
}
