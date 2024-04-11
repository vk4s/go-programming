package main

import (
	"fmt"
	"net/http"
	"os"
)

func init() {
	// `init` function runs when the module is initialized. That means it runs first.
	fmt.Println("Welcome to website health checker.")
}

func checkWebsiteStatus() {

	var website string
	fmt.Println("Enter the website address:")
	fmt.Scan(&website)

	response, err := http.Get(website)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if response.StatusCode >= 400 {
		fmt.Println("Website found but got status code", response.StatusCode)
		os.Exit(0)
	}

	fmt.Println("Website is online with status code", response.StatusCode)
}
