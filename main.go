package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	// URL to start crawling from
	startURL := "https://example.com"

	// Initialize a map to keep track of visited URLs
	visited := make(map[string]bool)

	// Create a channel to send URLs to be crawled
	queue := make(chan string)

	// Start the crawling process by visiting the initial URL
	go func() {
		queue <- startURL
	}()

	for url := range queue {
		// Mark the URL as visited
		visited[url] = true

		// Fetch the web page
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			continue
		}
		defer resp.Body.Close()

		// Parse the HTML content
		tokenizer := html.NewTokenizer(resp.Body)
		for {
			tokenType := tokenizer.Next()
			switch tokenType {
			case html.ErrorToken:
				return
			case html.StartTagToken, html.SelfClosingTagToken:
				token := tokenizer.Token()
				if token.Data == "a" {
					for _, attr := range token.Attr {
						if attr.Key == "href" {
							// Extract the link URL
							linkURL := attr.Val

							// Ensure the link URL is not visited already
							if !visited[linkURL] {
								fmt.Println(linkURL)
								// Add the link URL to the queue for further crawling
								go func() {
									queue <- linkURL
								}()
							}
						}
					}
				}
			}
		}
	}
}
