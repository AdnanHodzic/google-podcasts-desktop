package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

var newBaseUrl = "https://podcasts.google.com/?feed="

func parseURL(urlString string, field string) (string, error) {

	var result string

	url, err := url.Parse(urlString)
	if err != nil {
		return result, err
	}

	result = url.RawQuery
	// ToDo:Add check if not starting as proper Google Podcast URL
	if field != "" && result != "" {
		result = url.Query().Get(field)
	}
	return result, err
}

func usage() {
	appName := "--- Get Google Podcast web friendly URL's ---"
	fmt.Println("\n", appName+"\n")
	fmt.Println("Usage:")
	fmt.Println("  google-podcast-web-url https://www.google.com/podcast...")
}

func main() {
	flag.Usage = usage

	// ugly alternative
	//fieldPtr := flag.String("url=feed", "", "Google Podcast URL")
	fieldPtr := flag.String("url", "", "Google Podcast URL")

	flag.Parse()

	if *fieldPtr != "" {
		fmt.Println("\nSay something, empty!")
		fmt.Println(os.Args[0])
	}

	if len(flag.Args()) < 1 {
		//panic(errors.New("Please proper Google Podcast URL"))
		fmt.Println("\nError: You need to provide valid Google Podcast URL!")
		usage()
	} else {
		urlString := flag.Args()[0]

		resultURL, err := parseURL(urlString, *fieldPtr)
		if err != nil {
			panic(err)
			//fmt.Println("Don't panic")
		}

		fmt.Println("\nGoogle Podcast web friendly URL:")
		fmt.Println(newBaseUrl + resultURL)
	}
}
