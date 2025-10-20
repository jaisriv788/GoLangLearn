package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaisriv788/mongodb/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server Is Getting Started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000.")
}
