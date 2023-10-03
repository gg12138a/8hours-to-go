package main

import (
	"fmt"
	"time"
)

func sleepOrPrintTask(taskName string, t int) {
	for i := 0; i < t; i++ {
		fmt.Printf("task: %s\n", taskName)
		time.Sleep(1 * time.Second)
	}
}

func foo() bool {
	return true
}

func main() {
	go sleepOrPrintTask("a", 100)
	go sleepOrPrintTask("b", 100)

	for {
		fmt.Printf("main\n")
		time.Sleep(1 * time.Second)
	}

	// error, goroutine的通信需要使用channel
	// b := go foo()
}
