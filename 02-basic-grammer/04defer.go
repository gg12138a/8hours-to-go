package main

import "fmt"

func autoRelease() {
	x := 1

	// 清理资源
	defer fmt.Println("release, x = ", x)

	a := x * 2
	fmt.Println("a is calculated to be", a)
}

func returnBeforeDefer() int {
	defer func() {
		fmt.Println("defer...")
	}()

	x := func() int {
		fmt.Println("return...")
		return 1
	}

	return x() // 先return, 后defer
}

func main() {
	autoRelease()

	_ = returnBeforeDefer()
}
