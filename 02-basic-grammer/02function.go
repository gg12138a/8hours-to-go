package main

import "fmt"

// 定义函数的形参和返回值的类型
func functionParamType(key string, val int) int {
	fmt.Printf("%s = %d\n", key, val)
	return 2 * val
}

// 返回元组
func returnTuple() (string, int) {
	return "a", 2
}

// 在返回的元祖中，定义形参名
func returnNamedTuple() (rKey string, rVal int) {
	rKey = "a" // 重新赋值，默认值为""
	rVal = 20

	return
}

func returnSameTypeTuple() (a, b int) {
	return
}

func main() {
	_ = functionParamType("a", 1)

	tKey, tVal := returnTuple()
	fmt.Printf("tuple %s=%d\n", tKey, tVal)

	tKey, tVal = returnNamedTuple()
	fmt.Printf("tuple %s=%d\n", tKey, tVal)
}
