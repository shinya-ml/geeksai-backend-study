package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(_ http.ResponseWriter, _ *http.Request) {
		fmt.Println("pong")
	})

}
