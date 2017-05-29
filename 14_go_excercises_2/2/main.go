package main

import "fmt"

func main()  {
  add := func(n int) (int, bool)  {
    return n / 2, n%2 == 0
  }
  fmt.Println(add(123))
	fmt.Println(add(124))
}
