package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Now Listening on 5000...")
	fmt.Println("SEA_COLOR =", os.Getenv("SEA_COLOR"))
	http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}
