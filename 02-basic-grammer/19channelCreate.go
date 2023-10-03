package main

import "fmt"

func produce(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)
}

func main() {
	c := make(chan int)

	go produce(c)

	for i := range c {
		fmt.Println(i)
	}
}
