package main

import (
	"encoding/json"
	"os"
)

/*
  Encoding - writer
  Decoding - reader
*/

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"Iris", "West Allen", 23, 1}

	// NewEncoder gives *Encoder then we call Encode method with p1
	// func NewEncoder(w io.Writer) *Encoder  - NewEncoder syntax
	// func (enc *Encoder) Encode(v interface{}) error - Encode syntax
	// os.Stdout gives Writer so New Encoder get pointer to Encoder
	json.NewEncoder(os.Stdout).Encode(p1)
}
