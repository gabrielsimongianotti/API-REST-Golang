package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	fmt.Print("http://localhost:8080/")

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
