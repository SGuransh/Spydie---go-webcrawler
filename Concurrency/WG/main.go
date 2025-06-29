package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	wg.Add(1) // Add 1 to the WaitGroup counter

	msg = "Hello, World!"
	go updateMessage("Hello, universe!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)

	go updateMessage("Hello, cosmos!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)

	go updateMessage("Hello, world!", &wg)
	wg.Wait()
	printMessage()
}
