package main

import (
	"io"
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)

}