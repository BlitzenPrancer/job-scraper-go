package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/BlitzenPrancer/job-scraper"
)

const fileName string = "jobs.csv"

// Handler
func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)

	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	// fmt.Println(term)

	scrapper.Scrape(term)

	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/search", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
