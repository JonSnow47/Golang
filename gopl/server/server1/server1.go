// Server1 is my first web written by Go source code.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", server1)
	http.HandleFunc("/hello", handleHello)
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello, world!")
}

func server1(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "URL.Path: %s.", req.URL.Path)
}