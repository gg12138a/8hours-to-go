package main

import "fmt"

type AnimalIF interface {
	eat(food string)
}

type Cat struct {
	name string
}

func (c Cat) eat(food string) {
	fmt.Printf("cat %s is eating %s\n", c.name, food)
}

type Dog struct {
	name string
}

func (d Dog) eat(food string) {
	fmt.Printf("dog %s is eating %s\n", d.name, food)
}

func void(animal AnimalIF) {
	//todo
}

func main() {
	c := Cat{name: "li"}
	c.eat("pie")

	var i AnimalIF = c
	i.eat("pie")

	void(c)
}
