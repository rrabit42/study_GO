package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://job.incruit.com/entry/"

func main() {
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res := getHttp(pageURL)
	defer res.Body.Close() // res.Body는 byte인데 IO임(입력과 출력) -> 따라서 이 함수가 끝나면 닫아야함, 메모리 누출 막기
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title := card.Find(".title>a").Text()
		location := card.Find(".sjcl").Text()
		salary := card.Find(".salary>a").Text()
		summary := card.Find(".summary>a").Text()
	})
}

// http.Get(url) 하면 403
func getHttp(url string) *http.Response {
	fmt.Println("[GET] Request: ", url)
	req, rErr := http.NewRequest("GET", url, nil)
	checkErr(rErr)
	req.Header.Add("User-Agent", "Crawler")

	client := &http.Client{}
	res, err := client.Do(req)
	checkErr(err)
	checkCode(res)

	return res
}

func getPages() int {
	pages := 0

	res := getHttp(baseURL)
	defer res.Body.Close() // res.Body는 byte인데 IO임(입력과 출력) -> 따라서 이 함수가 끝나면 닫아야함, 메모리 누출 막기
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
