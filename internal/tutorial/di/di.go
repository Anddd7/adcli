package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(buffer io.Writer, s string) {
	fmt.Fprintf(buffer, "Hello, %s", s)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
