package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root")
	})
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
	router.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	router.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	router.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
	http.ListenAndServe(":8005", router) /// <------ (router)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	bookname := r.FormValue("name")
	fmt.Fprintf(w, "CreateBook : %s name : %s", title, bookname)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "ReadBook: %s \n", title)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	bookname := r.FormValue("name")
	fmt.Fprintf(w, "UpdateBook : %s name : %s", title, bookname)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	bookname := r.FormValue("name")
	fmt.Fprintf(w, "DeleteBook : %s name : %s", title, bookname)
}
