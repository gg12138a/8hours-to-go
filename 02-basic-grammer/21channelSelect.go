package main

import "fmt"

func fibonacci(nums chan int, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case nums <- x:
			// x写入nums成功
			x = y
			y = x + y
		case <-quit:
			// 从quit中获取一个值成功
			break
		}
	}
}

func main() {
	nums := make(chan int)
	quit := make(chan int)

	go fibonacci(nums, quit)

	for i := 0; i < 10; i++ {
		x := <-nums
		fmt.Println(x)
	}

	quit <- 0
}
