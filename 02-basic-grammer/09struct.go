package main

import "fmt"

type Book struct {
	title string
	auth  string
	price int
}

func (this *Book) calculateSum(cnt int) int {
	return this.price * cnt
}

func createInstance() Book {
	var b1 Book
	b1.title = "rust"
	b1.auth = "cool team"
	b1.price = 3

	fmt.Println(b1)

	// way 2:
	_ = Book{title: "rust", auth: "cool team", price: 3}

	return b1
}

func callMethod(b Book) {
	fmt.Println(b.calculateSum(2))
}

func main() {
	b := createInstance()
	callMethod(b)
}
