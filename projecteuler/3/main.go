package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lagestPrimeFactor(600851475143))
}

func lagestPrimeFactor(n int) int {
	var ans []int
	for n%2 == 0 {
		n /= 2
		ans = append(ans, 2)
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			n /= i
			ans = append(ans, i)
		}
	}
	if n > 2 {
		ans = append(ans, n)
	}
	return ans[len(ans)-1]
}

// Description
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
// Answer : 6857
