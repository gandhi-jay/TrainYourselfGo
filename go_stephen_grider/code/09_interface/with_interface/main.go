package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

type gujaratiBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	gb := gujaratiBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(gb)
}

func (englishBot) getGreeting() string {
	return "Hello, There"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (gujaratiBot) getGreeting() string {
	return "જય ગાંધી"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Concrete Type
// map
// struct
// int
// string
// englishBot

// Interface Type
// bot
