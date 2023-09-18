package main

import (
	"import-test/lib1"
	_ "import-test/lib2" // 让init函数执行
	. "import-test/lib3"
)

func main() {
	lib1.Bar()
	Qux()
}
