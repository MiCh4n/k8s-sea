package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func pirate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "There are pirates on the sea! Pods are in danger!")
}

func handleRequests() {
	LISTEN_ADDR := os.Getenv("LISTEN_ADDRESS")
	http.HandleFunc("/", pirate)
	log.Fatal(http.ListenAndServe(LISTEN_ADDR, nil))
}

func main() {
	handleRequests()
}
