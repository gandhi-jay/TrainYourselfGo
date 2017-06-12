package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"SpeedForceAge"`
}

func main() {
	var p1 person
	fmt.Println(p1) // uninitialize, default values

	bs := []byte(`{"first":"Black","last":"Flash","SpeedForceAge":23}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("-------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}

// For tags read - https://eager.io/blog/go-and-json/
