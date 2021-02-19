package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// panic은 컴파일러가 못찾아낸 error
	// 초기화된 map에 값을 추가할 수 없음 -> map이 nil이 되어버리기 때문
	// 따라서 빈 map을 선언하고 싶으면 뒤에 {} 빈 중괄호 넣어주기
	// var results = map[string]string{}
	// 혹은 make() : map을 만들어주는 func 사용, emtpy map을 초기화하고 싶을 때

	var results = make(map[string]string)

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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	// fmt.Println(results) 하면 map[key:value, key:value...] 이런식으로 나옴
	for url, result := range results {
		fmt.Println(url, result)
	}
}

// 웹사이트 접속
// go lang std library 이용
func hitURL(url string) error {
	// logger
	fmt.Println("Checking:", url)

	// request 보내고, response와 error 받기
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 { // 에러가 있다면!
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
