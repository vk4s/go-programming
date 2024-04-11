package main

import (
	"fmt"
	"net/http"
	"os"
)

/*
This function checks status of a website and returns the `status code`.

Parameters:
  - url string

Return:
  - statusCode int
  - error error
*/
func checkStatus(url string) (int, error) {

	resp, err := http.Head(url)

	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

func main() {
	websites := []string{
		"http://google.com",
		"http://wikipedia.org",
		"http://golang.org",
		"http://facebook.com",
		"http://amazon.com",
		"http://coursedl.org",
	}

	for _, website := range websites {
		statusCode, err := checkStatus(website)
		if err != nil {
			fmt.Println(website, "Error:", err)
			os.Exit(1)
		}
		fmt.Println(website, statusCode)
	}

}
