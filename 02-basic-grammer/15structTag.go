package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string `info:"nothing" version:"0"`
}

func main() {
	s := MyStruct{name: "abc"}

	ty := reflect.TypeOf(s)
	for i := 0; i < ty.NumField(); i++ {
		tag := ty.Field(i).Tag
		name := ty.Field(i).Name

		fmt.Printf("%s: %s\n", name, tag)
	}
}
