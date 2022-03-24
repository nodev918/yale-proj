package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

func main() {
	http.HandleFunc("/hi", sayHi)
	http.HandleFunc("/hello", sayHello)

	http.ListenAndServe(":3333", nil)
}
