package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	Greet(w, name)
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetHandler)))
}
