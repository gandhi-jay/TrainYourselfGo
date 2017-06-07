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

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Creating Map with key of int
	// value to another Map
	// which has key of string and value of int (count of occurances)
	//
	bucket := make(map[int]map[string]int)
	// Create slices to hold words words
	for i := 0; i < 12; i++ {
		bucket[i] = make(map[string]int)
	}
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		bucket[n][word]++
	}

	for k, v := range bucket[7] {
		fmt.Println(v, " \t- ", k)
	}
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
