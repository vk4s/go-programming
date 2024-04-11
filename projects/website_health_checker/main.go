package main

import (
	"fmt"
	"net/http"
)

func checkURL(url string, channel chan string) {
	_, err := http.Head(url)

	if err != nil {
		// fmt.Println(url, "might be down.")
		channel <- url + " might be down."
		return
	}
	// fmt.Println(url, "is up.")
	channel <- url + " is up."
}

func main() {
	websites := []string{
		"http://google.com",
		"http://wikipedia.org",
		"http://golang.org",
		"http://facebook.com",
		"http://amazon.com",
	}

	/* This channel will communicate between go routines via a `string` type of message
	Channels are bi-derectional way of communication.
	The operator <- (left angular bracket and a dash/hyphen) is used for sending messages.
	Send a value into a channel:
		channel <- "Hello"

	Receive the value from a channel:
		var message string
		message <- channel
	*/
	channel := make(chan string)

	for _, url := range websites {
		go checkURL(url, channel)
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-channel)
	}
}
