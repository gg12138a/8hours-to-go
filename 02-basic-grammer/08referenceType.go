package main

import "fmt"

func sliceTest() {
	var s1 []int = nil
	// 0x0, 0xc000010018
	fmt.Printf("%p, %p\n", s1, &s1)

	s1 = make([]int, 3)
	// 0xc00001a210, 0xc000010018
	fmt.Printf("%p, %p\n", s1, &s1)

	s2 := s1
	// 0xc00001a210, 0xc000010048
	fmt.Printf("%p, %p\n", s2, &s2)

	s3 := make([]int, 0)
	// 0x54e3e0, 0xc000010078
	fmt.Printf("%p, %p\n", s3, &s3)
}

func mapTest() {
	var m1 map[string]string = nil
	// 0x0, 0xc000050028
	fmt.Printf("%p, %p\n", m1, &m1)

	m1 = make(map[string]string)
	// 0xc00007a090, 0xc000050028
	fmt.Printf("%p, %p\n", m1, &m1)

	m2 := m1
	// 0xc00007a090, 0xc000050030
	fmt.Printf("%p, %p\n", m2, &m2)
}

func main() {
	sliceTest()

	fmt.Println()

	mapTest()
}
