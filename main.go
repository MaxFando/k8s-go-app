package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	port := "8080"
	log.Printf("start server on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Hello World!\n")
}
