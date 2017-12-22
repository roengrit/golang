package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root")
	})
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "product")
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "users")
	})
	http.ListenAndServe(":8005", nil)

}
