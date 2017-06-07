package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)

	scanner.Split(bufio.ScanWords)
	// Create slice of slice of string to hold slices of words
	bucket := make([][]string, 12)

	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		bucket[n] = append(bucket[n], word)
	}

	// Print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(bucket[i]))
	}
	// Print the words in one of the buckets
	// fmt.Println(bucket[7])
	fmt.Println(len(bucket))
	fmt.Println(cap(bucket))
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
