package main

import (
	"fmt"
	"log"
	"net/http"
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

//2ns server
// var mu sync.Mutex
// var count int

// func main() {
// 	http.HandleFunc("/", handler) //each request calls handler
// 	http.HandleFunc("/count", counter)
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }

// //handler echoes the path of the request URL
// func handler(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

// func counter(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	fmt.Fprintf(w, "Count %d\n", count)
// 	mu.Unlock()
// }

func main() {
	http.HandleFunc("/", handler) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handler echoes the path of the request URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote address = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q = %q", k, v)
	}
}
