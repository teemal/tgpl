package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//first server
// func main() {
// 	http.HandleFunc("/", handler) //each request calls handler
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }

// //handler echoes the path of the request URL
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) //each request calls handler
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handler echoes the path of the request URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
