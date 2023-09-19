package main

import "fmt"

// 只能传4[int]. 而且是按值拷贝的
func print4LenArr(arr [4]int) {
	for i, _ := range arr {
		arr[i] = 2 * arr[i]
	}

	for idx, val := range arr {
		fmt.Printf("arr[%d] = %d\n", idx, val)
	}
}

// slice按值传递，其内部具有三个属性(指针，len, capacity)
func printAnyLenArr(slice []int) {
	for i, _ := range slice {
		slice[i] = 2 * slice[i]
	}

	for idx, val := range slice {
		fmt.Printf("arr[%d] = %d\n", idx, val)
	}
}

func arrayOperate() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 10
	}

	arr2 := [10]int{12, 3}
	fmt.Println(arr2)

	for idx, val := range arr2 {
		fmt.Printf("arr2[%d] = %d\n", idx, val)
	}
}

func sliceOperate() {
	slice := []int{1, 2, 3}
	fmt.Printf("%T\n", slice) // []int

	printAnyLenArr(slice)
	fmt.Println(slice)
}

func createSlice() {
	var slice []int // slice内部的指针为nil
	// slice[0] = 10 // error

	slice = make([]int, 1)
	slice[0] = 10
}

func sliceNil() {
	var s []int
	fmt.Println(len(s)) // nil slice的长度为0

	var slice1 []int // 声明但未初始化
	slice2 := []int{}

	fmt.Println(slice1 == nil) // true
	fmt.Println(slice2 == nil) // false

}

func main() {
	arrayOperate()

	arr := [4]int{1, 2, 3}
	print4LenArr(arr)
	//print4LenArr([3]int{1, 2, 3}) error，类型不匹配
	fmt.Println(arr) // 由于是值传递的，因此func内的修改无效

	sliceOperate()

	printAnyLenArr(arr[:])

	createSlice()

	sliceNil()
}
