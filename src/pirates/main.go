package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func pirate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "There are pirates on the sea! Pods are in danger!\n")
}

func handleRequests() {
	Host := os.Getenv("HOST_PIRATE")
	Port := os.Getenv("PORT_PIRATE")
	http.HandleFunc("/pirate", pirate)
	log.Fatal(http.ListenAndServe(Host+":"+Port, nil))
}

func main() {
	handleRequests()
}
