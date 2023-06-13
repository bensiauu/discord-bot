package collector

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Collector interface {
	Visit(URL string) error
}

type collector struct {
	scraper *colly.Collector
}

func (c *collector) Visit(url string) error {
	return c.scraper.Visit(url)
}

func NewCollector() Collector { //possibly expand in this file if need to add in opts for collector
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	return &collector{
		scraper: c,
	}
}
