package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

var signals = []string{"test"}

func main() {
	// go greeting("Hello")
	// greeting("world")
	websitelist := []string{
		"https://go.dev", "https://google.com", "https://fb.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeting(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(1 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		log.Fatal("Something Went Wrong.")
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d Status Code for %s\n", res.StatusCode, endpoint)
}
