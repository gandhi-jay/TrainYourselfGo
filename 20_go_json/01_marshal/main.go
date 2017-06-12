package main

import (
	"encoding/json"
	"fmt"
)

// First letter should be capital to Marshal into JSON
type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

// All not Exported.
type speedster struct {
	name  string
	speed int
}

func main() {
	p1 := person{"Herisson", "Wells", 40, 87}
	bs, _ := json.Marshal(p1) // from encoding/json package
	fmt.Println(bs)           // Slice of ints which is ascii values of runes
	fmt.Printf("%T\n", bs)    // []uint8
	fmt.Println(string(bs))   // JSON Data

	p2 := speedster{"Flash", 670}
	p22js, _ := json.Marshal(p2)
	fmt.Println(p22js)         // Empty Slice of ints which is ascii values of runes
	fmt.Printf("%T\n", p22js)  // []uint8
	fmt.Println(string(p22js)) // Empty JSON Data, only braces.
}

/*
	encoding
		Marshal/Unmarshal - json to string/string to json
		encode/decode - stream

*/
