package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello world...")
		return
	})

	log.Println(http.ListenAndServe(":8080", nil))
}
