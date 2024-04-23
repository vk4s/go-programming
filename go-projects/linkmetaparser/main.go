package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type LinkPreview struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func getURLMeta(url string) LinkPreview {
	resp, err := http.Get(url)
	if err != nil {
		return LinkPreview{}
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return LinkPreview{}
	}

	var title, description, image string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			title = strings.TrimSpace(n.FirstChild.Data)
		}
		if n.Type == html.ElementNode && n.Data == "meta" {
			for _, attr := range n.Attr {
				if attr.Key == "name" && attr.Val == "description" {
					for _, a := range n.Attr {
						if a.Key == "content" {
							description = a.Val
						}
					}
				}
				if attr.Key == "property" && attr.Val == "og:image" {
					for _, a := range n.Attr {
						if a.Key == "content" {
							image = a.Val
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return LinkPreview{
		Title:       title,
		Description: description,
		Image:       image,
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	urls := []string{"https://lorbic.com", "https://google.com", "https://instagram.com", "https://r.lorbic.com"} // List of URLs
	var linkPreviews []LinkPreview

	for _, url := range urls {
		linkPreviews = append(linkPreviews, getURLMeta(url))
	}

	jsonResponse, err := json.Marshal(linkPreviews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println("Server running on port 8081")
	done := make(chan bool)
	go func() {
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	<-done
}
