package main

import (
	"fmt"
	"init-and-main-fn/lib1"
	_ "init-and-main-fn/lib2" // 副作用导入
)

func main() {
	fmt.Println("lib1.Username = ", lib1.UserName)
}
