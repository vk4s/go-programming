package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkURL(url string, channel chan string) {
	_, err := http.Head(url)

	if err != nil {
		fmt.Println(url, "might be down.")
		channel <- url
		return
	}
	fmt.Println(url, "is up.")
	channel <- url
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

	// repeating routines
	// Wait for the channel to send a new url and when it is received run the body of for loop
	for url := range channel {
		go func(_url string) {
			// It reruns the request after 3 seconds
			time.Sleep(2 * time.Second)
			checkURL(_url, channel)
		}(url)
	}
}
