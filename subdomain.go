package main

import (
	"log"

	"github.com/gocolly/colly"
)

type SubdomainScraperProvider[T SubdomainResult] struct{}

func (s *SubdomainScraperProvider[SubdomainResult]) Scrape(subd string) SubdomainResult {
	items := []SubdomainItem{}
	col := colly.NewCollector()

	col.OnHTML("table#table > tbody > tr", func(ele *colly.HTMLElement) {
		ch := ele.DOM.Children()
		domain := ch.Eq(1).Text()
		address := ch.Eq(2).Children().Eq(0).Text()
		rtype := ch.Eq(3).Text()
		date := ch.Eq(4).Text()

		items = append(items, SubdomainItem{domain, address, rtype, date})
	})

	url := "https://rapiddns.io/subdomain/" + subd + "?full=1"
	err := col.Visit(url)

	if err != nil {
		log.Fatalln(err)
	}

	return SubdomainResult{items}
}

type SubdomainResult struct {
	Items []SubdomainItem `json:"items"`
}

type SubdomainItem struct {
	Domain     string `json:"domain"`
	Address    string `json:"address"`
	RecordType string `json:"record_type"`
	Date       string `json:"date"`
}

type SubdomainClient struct {
	Domain *string
	Result *SubdomainResult
	S      Scraper[SubdomainResult]
}

func (c *SubdomainClient) Search() SubdomainResult {
	if c.Domain == nil {
		log.Fatal("subdomain is not specified.")
	}

	return c.S.Scrape(*c.Domain)
}
