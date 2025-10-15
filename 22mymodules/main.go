package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the server")
	greeter()

	r := mux.NewRouter()

	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func greeter() {
	fmt.Println("Hey their users.")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}
