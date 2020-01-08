package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	msg := "Hello World Server🐱"
	http.Handle("/", String(msg))
	http.ListenAndServe(":9998", nil)
}
