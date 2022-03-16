package generator

import (
	"math/rand"
	"sync"
	"time"
)

//Generator записывает новые экзмепляры генератора
type Generator struct {
	limit     int
	inCommand chan string
}

// NewGenerator возвращает ссылку на новый экземпляр генератора случайных чисел и канал передачи комманд
func NewGenerator(limit int) (*Generator, chan string) {
	g := &Generator{limit: limit, inCommand: make(chan string, 1)}
	return g, g.inCommand
}

// RandomNumber запускает рабочий цикл генератора случайных чисел
func (g *Generator) RandomNumber(wg *sync.WaitGroup, number chan int) {
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case stopCommand := <-g.inCommand:
			if stopCommand == "stop" {
				wg.Done()
				return
			}
		default:
			if len(number) < g.limit {
				randomNumber := rand.Intn(g.limit + 1)
				number <- randomNumber
			}
		}
	}
}
