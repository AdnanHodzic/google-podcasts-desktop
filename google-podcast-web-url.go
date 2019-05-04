package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

var newBaseUrl = "https://podcasts.google.com/?"

func parseURL(urlString string) (string, error) {

	var result string

	url, err := url.Parse(urlString)
	if err != nil {
		return result, err
	}

	result = url.RawQuery
	return result, err
}

func main() {

	fmt.Println("\nEnter Google Podcast URL:")
	reader := bufio.NewReader(os.Stdin)
	providedURL, _ := reader.ReadString('\n')

	resultURL, err := parseURL(providedURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nGoogle Podcast web friendly URL:")
	fmt.Println(newBaseUrl + resultURL)
}
