package main

import (
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=golang&start=0"

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

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("Something went wrong while closing response body")
		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination")

	return 0
}

func main() {
	getPages()
}
