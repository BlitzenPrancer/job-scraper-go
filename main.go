package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type jobItem struct {
	id       string
	title    string
	company  string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://www.indeed.com/jobs?limit=50&l=Seattle&q=software+engineer"

func main() {
	var jobs []jobItem
	pages := getNumOfPages()

	for i := 0; i < pages; i++ {
		jobItemsOnPage := getPage(i)
		jobs = append(jobs, jobItemsOnPage...)
	}

	fmt.Println(jobs[0])

}

// getPage parses the page and returns jobs on that page
func getPage(page int) []jobItem {
	var jobsOnPage []jobItem
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)

	res, err := http.Get(pageURL)
	CheckError(err)
	defer res.Body.Close()
	CheckCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckError(err)
	jobCards := doc.Find(".jobsearch-SerpJobCard")
	jobCards.Each(func(i int, card *goquery.Selection) {
		job := initJobItem(card)
		jobsOnPage = append(jobsOnPage, job)
	})

	return jobsOnPage
}

func initJobItem(card *goquery.Selection) jobItem {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title > a").Text())
	company := CleanString(card.Find(".company > a").Text())
	location := CleanString(card.Find(".location").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())

	return jobItem{id: id, title: title, company: company, location: location, salary: salary, summary: summary}
}

// getNumOfPages gets total number of pages from the result.
func getNumOfPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	CheckError(err)
	defer res.Body.Close()
	CheckCode(res)

	// Load HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	pages = doc.Find(".pagination-list a").Length()
	return pages
}
