package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")
	done <- true // Signal completion
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done // Wait for completion
}
