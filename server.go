package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "<h1> Hello, I'm %s and I'm %s years old!!! </h1>", name, age) //w.Write([]byte("<h1> Hello, FullCycle!!! </h1>"))
}
