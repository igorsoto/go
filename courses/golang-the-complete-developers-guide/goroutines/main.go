package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}
	
	c := make(chan string)
	
	for _, url := range urls {
		go healthCheck(url, c)
	}

	for u := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			healthCheck(url, c)
		}(u)
	}
}

func healthCheck(url string, c chan string) {
	_, err := http.Head(url)
	if err != nil {
		fmt.Println(url, "is down")
		c <- url
		return
	}
	fmt.Println(url, "is up")
	c <- url
}