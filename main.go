package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	ERR_GET_HTTP = 1
	ERR_IO_READ  = 2
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		os.Exit(ERR_GET_HTTP)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		os.Exit(ERR_IO_READ)
	}

	fmt.Println(string(body))
}
