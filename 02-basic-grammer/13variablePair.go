package main

import "fmt"

func main() {
	// pair<statictype:string, value:"abc">
	var s string = "abc"

	// pair<statictype:string, value:"abc">
	var i interface{} = s

	val, ok := i.(string)
	if ok {
		fmt.Println(val)
	}
}
