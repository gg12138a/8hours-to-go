package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) eat(food string) {
	fmt.Printf("%s is eating %s\n", this.name, food)
}

type SuperMan struct {
	Human // 继承
	power int
}

func (this *SuperMan) fly() {
	fmt.Printf("%s is flying\n", this.name)
}

func main() {
	{
		h := Human{"abc", "f"}
		h.eat("apple")
	}

	{
		s := SuperMan{Human{"abc", "f"}, 18}
		s.fly()
	}

}
