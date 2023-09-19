package main

import "fmt"

func logSlice(s []int) {
	if s == nil { // 比较的是slice中的[addr]
		fmt.Printf("slice addr = %p, len = %d, cap = %d\n", &s, len(s), cap(s))
		return
	}

	fmt.Printf("slice addr = %p, first elem addr = %p, len = %d, cap = %d\n", &s, &s[0], len(s), cap(s))
}

func sliceAppend() {
	// slice addr = 0xc000010018, first elem addr = 0xc0000200c0, len = 3, cap = 5
	s := make([]int, 3, 5)

	// slice addr = 0xc000010030, first elem addr = 0xc0000200c0, len = 4, cap = 5
	a := append(s, 1)

	var n []int = nil
	// slice addr = 0xc000010060, first elem addr = 0xc0000120d8, len = 1, cap = 1
	n = append(n, 1)

	logSlice(s)
	logSlice(a)
	logSlice(n)
}

func sliceCopy() {
	// slice addr = 0xc000010078, first elem addr = 0xc0000200f0, len = 3, cap = 5
	ori := make([]int, 3, 5)

	// slice addr = 0xc000010090, first elem addr = 0xc0000200f0, len = 3, cap = 5
	lowCopy := ori

	// slice addr = 0xc000010090, first elem addr = 0xc00001a1f8, len = 3, cap = 3
	dst := make([]int, len(ori))
	copy(dst, ori)

	logSlice(lowCopy)
	logSlice(ori)
	logSlice(dst)
}

func main() {
	logSlice(nil)

	sliceAppend()

	fmt.Println()

	sliceCopy()
}
