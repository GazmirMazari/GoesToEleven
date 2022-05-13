package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)

		return
	}

	if r.Method != "GET" {
		http.Error(w, "Required method is Get", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)

		return
	}

	if r.Method != "POST" {
		http.Error(w, "Required method is POST", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Form request was successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 5000\n")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}

}
