package main

import (
	"fmt"
	"time"
)

func main() {
	// go start a goroutine
	for i := 0; i < 5; i++ {
		go printHelloWorld(i)
	}
	time.Sleep(5 * time.Millisecond)
}

func printHelloWorld(i int) {
	fmt.Printf("Hello, World!  %d\n", i)
}
