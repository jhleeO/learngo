package main

import (
	"errors"
	"fmt"
	"net/http")


var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.naver.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.instagram.com",
		"https://www.google.com",
		"https://www.reddit.com",
		"https://www.github.com",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}