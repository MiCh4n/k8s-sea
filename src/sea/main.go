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
	HostS := os.Getenv("HOST_SEA")
	PortS := os.Getenv("PORT_SEA")
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, _Request)
	})

	log.Fatal(http.ListenAndServe(HostS+":"+PortS, nil))
}
func main() {
	HostP := os.Getenv("HOST_PIRATE")
	PortP := os.Getenv("PORT_PIRATE")
	response, err := http.Get("http://" + HostP + ":" + PortP)

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
