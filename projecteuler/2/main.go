package main

import "fmt"

func main() {
	var total int
	for i := 0; ; i++ {
		n := fibonacci(i)
		if n > 4000000 {
			break
		}
		if n%2 == 0 {
			total += n
		}
	}
	fmt.Println(total)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
