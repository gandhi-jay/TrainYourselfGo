package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person
	fmt.Println(p1) // uninitialize, default values

	bs := []byte(`{"First":"Black","Last":"Flash","Age":23}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("-------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
