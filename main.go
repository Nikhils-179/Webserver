package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not found", http.StatusNotFound)
	}

	fmt.Fprint(w, "hello Golang")
}

func formHandller(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error : %v", err)
	}
	fmt.Fprintln(w, "POST REQUEST SUCCESSFUL")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name : %s \n", name)
	fmt.Fprintf(w, "Address : %s \n", address)
	fmt.Println("Form Handler Called and exited")
}
func main() {
	fmt.Println("Starting the Server at localost:8080")
	fileserver := http.FileServer(http.Dir("./static")) // loads .index.html
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandller)
	http.HandleFunc("/hello", helloHandler)
	hostPort := "localhost:8080"
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started at port at port *8080")
	fmt.Println("Good Bye")
}
