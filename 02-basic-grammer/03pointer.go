package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {
	a, b := 1, 2
	swap(&a, &b)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
