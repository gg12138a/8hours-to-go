package main

import (
	"fmt"
)

func declareVariable() {
	var x int //默认为0
	fmt.Println("x =", x)

	var y int = 100
	fmt.Println("y =", y)

	var t = 100
	fmt.Printf("t = %d, type = %T\n", t, t)

	a := 20
	fmt.Printf("a = %d, type = %T\n", a, a)

	var v1, v2 int = 100, 200
	fmt.Println(v1, v2)

	var a1, a2 = 1, "a"
	fmt.Println(a1, a2)

	var (
		b1 = 100
		b2 = 200
	)
	fmt.Println(b1, b2)
}

func constVariable() {
	const length = 10

	const (
		BeiJing  = 0
		ShangHai = 1
		ShenZhen = 2
	)

	// iota从0开始计数
	const (
		BeijingA  = iota
		ShangHaiA = iota
		ShenZhenA = iota
	)

	// 若省略初始化表达式，则与上一行一致
	const (
		T1 = 1
		T2
		T3
	)

	// iota默认值为0, 每行递增
	const (
		A, B = iota + 1, iota + 2 // 1,2
		C    = iota               // 1
		D                         // 2
	)
}

func main() {
	declareVariable()
	constVariable()
}
