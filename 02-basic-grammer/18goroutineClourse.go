package main

import (
	"fmt"
	"time"
)

func main() {
	go func(s string) {
		fmt.Printf("task: %s\n", s)
	}("abc")

	time.Sleep(time.Second)
}
