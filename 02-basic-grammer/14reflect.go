package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age  int
	Name string
}

func reflectPerson(p Person) {
	ty := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < ty.NumField(); i++ {
		fieldTy := ty.Field(i)
		fieldVal := v.Field(i).Interface()

		// age int = %!s(int=18)
		// name string = abc
		fmt.Printf("%s %s = %s\n", fieldTy.Name, fieldTy.Type, fieldVal)
	}
}

func main() {
	p := Person{Age: 18, Name: "abc"}

	reflectPerson(p)
}
