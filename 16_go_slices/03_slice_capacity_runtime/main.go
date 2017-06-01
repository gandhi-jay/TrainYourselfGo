package main

import "fmt"

func main() {
	// Slice is reference type like map and channels
	// Slice points to underlying array.

	mySlice := make([]int, 0, 5)
	fmt.Println("--------------------------")
	fmt.Println(mySlice, "\t", len(mySlice), "\t", cap(mySlice))
	fmt.Println("--------------------------")

	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len", len(mySlice), "Cap", cap(mySlice), "Val", mySlice[i])
	}
	// When slice reaches it's thresold(beyond capacity), it doubles it's capacity.
	// Performance cost will come into play.

}
