package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetch(url string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	res, _ := http.Get(url)
	body, _ := io.ReadAll(res.Body)
	fmt.Printf("%s\n", body)
	ch <- url

}
func main() {

	urls := []string{"https://golang.org", "https://go.dev"}
	var wg sync.WaitGroup
	ch := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		fetch(url, &wg, ch)
	}
	wg.Wait()
	for i := range ch {
		fmt.Printf("Fetched: %s\n", i)
	}

}
