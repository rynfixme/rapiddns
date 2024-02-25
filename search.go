package main

import (
	"log"
	"net/url"

	"github.com/gocolly/colly"
)

type SearchScraperProvider[T SearchResult] struct{}

func (s *SearchScraperProvider[SearchResult]) Scrape(word string) SearchResult {
	items := []SearchItem{}
	col := colly.NewCollector()

	col.OnHTML("table#table > tbody > tr", func(ele *colly.HTMLElement) {
		ch := ele.DOM.Children()
		domain := ch.Eq(1).Text()
		address := ch.Eq(2).Children().Eq(0).Text()
		rtype := ch.Eq(3).Text()
		date := ch.Eq(4).Text()

		items = append(items, SearchItem{domain, address, rtype, date})
	})

	url := "https://rapiddns.io/s/" + url.QueryEscape(word) + "?full=1"
	err := col.Visit(url)

	if err != nil {
		log.Fatalln(err)
	}

	return SearchResult{items}
}

type SearchResult struct {
	Items []SearchItem `json:"items"`
}

type SearchItem struct {
	Domain     string `json:"domain"`
	Address    string `json:"address"`
	RecordType string `json:"record_type"`
	Date       string `json:"date"`
}

type SearchClient struct {
	Word   *string
	Result *SearchResult
	S      Scraper[SearchResult]
}

func (c *SearchClient) Search() SearchResult {
	if c.Word == nil {
		log.Fatalln("Search word is not specified.")
	}

	return c.S.Scrape(*c.Word)
}
