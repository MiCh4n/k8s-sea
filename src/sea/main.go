package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Request struct {
	Response string
	Color    bool
}

func page(data string, pirates bool) {
	_Request := Request{
		Response: data,
		Color:    pirates}
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, _Request)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
func main() {
	Host := os.Getenv("HOST_SEA")
	Port := os.Getenv("PORT_SEA")
	response, err := http.Get("http://" + Host + ":" + Port + "/pirate")

	if err != nil {
		c := true
		page("There are no pirates on the sea! Pods are safe!", c)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := false
	page(string(responseData), c)

}
