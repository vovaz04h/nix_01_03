package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))
}
