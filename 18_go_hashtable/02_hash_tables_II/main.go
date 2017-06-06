package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split function for the scanning
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts
	buckets := make([]int, 200)

	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:122])

}

func HashBucket(word string) int {
	return int(word[0])
}
