package main

import "fmt"

func main() {
	a := 32
	b := "GoLang"
	c := true
	d := `DD`
	e := 123.12

	fmt.Printf("%v \t %T \n",a,a)
	fmt.Printf("%v \t %T \n",b,b)
	fmt.Printf("%v \t %T \n",c,c)
	fmt.Printf("%v \t %T \n",d,d)
	fmt.Printf("%v \t %T \n",e,e)
}