package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1) // This is available to me from wg_challenge.go as its the same package
	updateMessage("Hello, universe!", &wg)
	wg.Wait()
	if msg != "Hello, universe!" {
		t.Errorf("Expected message to be 'Hello, universe!', but got '%s'", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	wg.Add(1)
	go updateMessage("Yo wassup", &wg)
	wg.Wait()
	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	result_str := string(result)
	os.Stdout = stdout
	if strings.TrimSpace(result_str) != "Yo wassup" {
		t.Errorf("Expected the message to be 'Yo wassup' but got this '%s'", result_str)
	}
}

func Test_main(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	result_str := string(result)
	os.Stdout = stdout
	if strings.TrimSpace(result_str) != "Hello, universe!\nHello, cosmos!\nHello, world!" {
		t.Errorf("Just wrong")
	}
}
