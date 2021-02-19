package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		hitURL(url)
	}
}

// 웹사이트 접속
// go lang std library 이용
func hitURL(url string) error {
	// logger
	fmt.Println("Checking:", url)

	// request 보내고, response와 error 받기
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
