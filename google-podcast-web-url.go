package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"
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

func openBrowser(linkURL string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", linkURL).Start()
	case "darwin":
		err = exec.Command("open", linkURL).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", linkURL).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	fmt.Println("\nEnter Google Podcasts URL:")
	reader := bufio.NewReader(os.Stdin)
	providedURL, _ := reader.ReadString('\n')

	if providedURL == "\n" {
		fmt.Println("You must enter a value.")
		os.Exit(0)
	}

	resultURL, err := parseURL(providedURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nGoogle Podcasts web friendly URL:")
	fmt.Println(newBaseUrl + resultURL)

	fmt.Println("\nOpening link in a default web browser ...")
	linkURL := newBaseUrl + resultURL
	openBrowser(linkURL)

}
