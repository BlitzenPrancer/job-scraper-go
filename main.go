package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.indeed.com/jobs?limit=50&l=Seattle&q=software+engineer"

func main() {
	pages := getNumOfPages()
	fmt.Println(pages)
}

// getNumOfPages gets total number of pages from the result.
func getNumOfPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code err: %d %s", res.StatusCode, res.Status)
	}

	// Load HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	pages = doc.Find(".pagination-list a").Length()
	return pages
}
