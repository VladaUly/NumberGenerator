package generator

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Generator записывает новые экзмепляры генератора
type Generator struct {
	limit        int
	inCommand    chan string
	startCommand chan string
}

// NewGenerator возвращает ссылку на новый экземпляр генератора случайных чисел и канал передачи комманд
func NewGenerator(limit int) (*Generator, chan string, chan string) {
	g := &Generator{limit: limit, inCommand: make(chan string, 1), startCommand: make(chan string)}
	return g, g.inCommand, g.startCommand
}

// RandomNumber запускает рабочий цикл генератора случайных чисел
func (g *Generator) RandomNumber(wg *sync.WaitGroup, number chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case stopCommand := <-g.inCommand:
			if stopCommand == "stop" {
				wg.Done()
				return
			}
		case stCommand := <-g.startCommand:
			if stCommand == "start" {
				fmt.Println("works")
				randomNumber := rand.Intn(g.limit + 1)
				number <- randomNumber
			}
		}
	}
}
