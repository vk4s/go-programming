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

	// - This line prints the request body to the terminal
	io.Copy(os.Stdout, resp.Body)

	// - This also prints the request body to the terminal
	// body, err := io.ReadAll(resp.Body)
	// resp.Body.Close()

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(len(string(body)))

}
