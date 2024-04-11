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
	// It reruns the request after 3 seconds
	time.Sleep(3 * time.Second)
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
	// it runs infinitly and starts back up when request is complete
	for {
		// the `<-channel` is a blocking statement
		// and waits for the request to complete before calling the function
		go checkURL(<-channel, channel)
	}
}
