package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Syntax
// func receiver_ func_name(args) return_type
// It is another parameter.
// Any variable with type person can call this func.
func (p person) fullname() string {
	return fmt.Sprint(p.first, " ", p.last)
}

// Notes from Todd
// (p person) is the "receiver"
// it is another parameter
// not idiomatic to use "this" or "self"
//
// "Not many people know this, but method notation, i.e. v.Method() is actually syntactic sugar and Go also understands the de-sugared version of it: (T).Method(v). You can see an example here. Naming the receiver like any other parameter reflects that it is, in fact, just another parameter quite well.
// This also implies that the receiver-argument inside a method may be nil. This is not the case with this in e.g. Java."
// SOURCE:
// https://www.reddit.com/r/golang/comments/3qoo36/question_why_is_self_or_this_not_considered_a/?utm_source=golangweekly&utm_medium=email

func main() {
	p1 := person{"Bruce", "Wayne", 28}
	p2 := person{"Barry", "Allen", 21}
	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
