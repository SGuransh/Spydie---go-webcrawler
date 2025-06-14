package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var msg string

func updateMesage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMesage("Hello, universe!", &mutex)
	go updateMesage("Hello, cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)

}
