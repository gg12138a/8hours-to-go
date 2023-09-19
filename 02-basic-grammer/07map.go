package main

import (
	"fmt"
)

func defineMap() {
	var m1 map[string]int
	if m1 == nil {
		println("nil")
	}
	fmt.Println(m1)

	m2 := make(map[string]string, 3)
	m2["name"] = "abc"
	m2["age"] = "18"
	fmt.Println(m2)

	m3 := map[string]string{
		"name": "abc",
		"age":  "18",
	}
	fmt.Println(m3)
}

func operateMap() {
	m := make(map[string]string)

	// add
	m["name"] = "abc"

	// iterate
	for k, v := range m {
		fmt.Printf("[%s] = %s\n", k, v)
	}

	// remove
	delete(m, "name")
}

func main() {
	defineMap()
	operateMap()

	m := make(map[string]int)
	fmt.Printf("%p, %p\n", m, &m)
}
