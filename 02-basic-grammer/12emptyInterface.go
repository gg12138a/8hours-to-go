package main

import "fmt"

func foo(arg interface{}) {
	fmt.Println(arg)

	// 类型断言
	val, ok := arg.(string)
	if ok {
		fmt.Printf("type is string, val is %s", val)
	}
}

type Bar struct {
	x int
}

type Log interface {
	log()
}

func (b *Bar) log() {
	fmt.Println(b)
}

func main() {
	foo(1)
	foo("abc")
	foo([3]int{1, 2, 3})

	foo(Bar{x: 1}) // {1}

	var l Log = &Bar{x: 1}
	foo(l) // &{1}
}
