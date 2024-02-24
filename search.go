package main

import (
	"log"

	"github.com/gocolly/colly"
)

type SearchScraperProvider struct{}

func (s *SearchScraperProvider) Scrape(word string) []SearchItem {
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

	url := "https://rapiddns.io/s/" + word + "?full=1"
	err := col.Visit(url)

	if err != nil {
		log.Fatalln(err)
	}

	return items
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
	Word   string
	Result SearchResult
	S      SearchScraperProvider
}

func (sc *SearchClient) search() error {
	domains := sc.S.Scrape(sc.Word)
	sc.Result.Items = domains
	return nil
}
