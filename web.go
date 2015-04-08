package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port   = os.Getenv("PORT")
	target = os.Getenv("GREETING")
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", target)
}

func main() {
	if port == "" {
		port = "3000"
	}
	if target == "" {
		target = "World"
	}

	http.HandleFunc("/", helloHandler)

	s := &http.Server{
		Addr: ":" + port,
	}
	log.Fatal(s.ListenAndServe())
}
