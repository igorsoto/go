package main

import (
	"log"
	"net/http"
)

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func main() {

	// using http.Handle
	http.Handle("/foo", &fooHandler{
		Message: "hello world",
	})

	// using http.HandleFunc
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello world 2"))
	})

	// start listening to requests, log and exit if it fails
	log.Fatal(http.ListenAndServe(":5000", nil))
}
