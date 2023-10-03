package main

import (
	"encoding/json"
	"fmt"
)

type MyJson struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []string `json:"friends"`
}

func main() {
	var j = MyJson{
		Name:    "abc",
		Age:     18,
		Friends: []string{"xiaoming", "xiaohong"},
	}

	jsonStr, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s\n", jsonStr)
	}

	fromJson := MyJson{}
	err = json.Unmarshal(jsonStr, &fromJson)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v\n", fromJson)
	}
}
