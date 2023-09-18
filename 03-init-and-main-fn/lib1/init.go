package lib1

import "fmt"

var (
	UserName string
)

func init() {
	fmt.Println("init lib1...")
	UserName = "abc"
}
