package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	url := "https://lorbic.com"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp.Header)

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(len(string(body)))

}
