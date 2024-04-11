package main

import (
	"fmt"
	"net/http"
)

func checkURL(url string) {
	_, err := http.Head(url)

	if err != nil {
		fmt.Println(url, "might be down.")
		return
	}
	fmt.Println(url, "is up.")
}

func main() {
	websites := []string{
		"http://google.com",
		"http://wikipedia.org",
		"http://golang.org",
		"http://facebook.com",
		"http://amazon.com",
	}

	for _, url := range websites {
		go checkURL(url)
	}

}
