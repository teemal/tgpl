//fetch all urls concurrently and report times and sizes
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start go routine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//main part of book
// func fetch(url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprint(err) //send to channel ch
// 		return
// 	}
// 	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
// 	resp.Body.Close() //dont leak resource
// 	if err != nil {
// 		ch <- fmt.Sprintf("while reading %s: %v", url, err)
// 		return
// 	}
// 	sec := time.Since(start).Seconds
// 	ch <- fmt.Sprintf("%.2fs   %7d   %s", sec, nbytes, url)
// }

//exercise 1.10
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close() //dont leak resource
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	sec := time.Since(start).Seconds
	ch <- fmt.Sprintf("%.2fs   %7d   %s", sec, nbytes, url)
}
