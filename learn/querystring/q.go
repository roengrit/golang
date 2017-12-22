package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		first := queryValues.Get("first") //work
		last := queryValues.Get("last")   //work
		fmt.Fprintf(w, first+"-->")
		fmt.Fprintf(w, last)
	})
	http.ListenAndServe(":8005", nil)

}
