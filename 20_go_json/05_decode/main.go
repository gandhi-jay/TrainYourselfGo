package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 person

	// func NewReader(s string) *Reader
	// takes string and gives point to Reader
	// NewReader returns a new Reader reading from s.
	// It is similar to bytes.NewBufferString but more efficient and read-only.
	rdr := strings.NewReader(`{"First":"Black","Last":"Flash","Age":23}`)
	fmt.Printf("Type of Reader : %T \n", rdr)

	// func NewDecoder(r io.Reader) *Decoder
	// takes Reader and gives pointer to Decoder
	// func (dec *Decoder) Decode(v interface{}) error
	// takes point to Decoder.
	// Decode reads the next JSON-encoded value from its input and
	// stores it in the value pointed to by v.
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println("-------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
