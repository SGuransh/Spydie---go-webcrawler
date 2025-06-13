package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4) // Add 4 to the WaitGroup counter

	go printSomething("Hello, World!", &wg)
	go printSomething("This is a simple Go program.", &wg)
	go printSomething("It prints messages to the console.", &wg)
	go printSomething("Goodbye!", &wg)

	wg.Wait()
}
