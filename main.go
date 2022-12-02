package main

import (
	"net/http"

	"github.com/1234567-kien/homecontroller"
)

func main() {
	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/home", homecontroller.Index)
	http.HandleFunc("/home/index", homecontroller.Index)

	http.ListenAndServe(":3000", nil)
}
