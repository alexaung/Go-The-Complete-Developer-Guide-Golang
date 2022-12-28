package main

import (
	"fmt"
	"net/http"
	"os"
)


func main()  {
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}