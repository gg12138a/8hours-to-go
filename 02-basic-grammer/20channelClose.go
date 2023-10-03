package main

import "fmt"

func main() {
	c := make(chan int)

	close(c)

	x, ok := <-c
	if ok {
		fmt.Println(x)
	}
}
