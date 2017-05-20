package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main()  {
	// No matter what, you've to use a & b.
	a := "Jay"
	b := "Gandhi"
	fmt.Println(a,b)

	// Using http code without blank identifier
	res, err := http.Get("http://www.discoverdollar.com")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	fmt.Printf("%s", page)

	// Using blank identifier
	res_, _ := http.Get("http://www.reportgarden.com")
	page_, _ := ioutil.ReadAll(res_.Body)
	res.Body.Close()
	fmt.Printf("%s", page_)
}
