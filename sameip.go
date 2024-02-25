package main

import (
	"log"

	"github.com/gocolly/colly"
)

type SameIPScraperProvider[T SameIpResult] struct{}

func (s *SameIPScraperProvider[SameIpResult]) Scrape(address string) SameIpResult {
	items := []SameIPItem{}
	col := colly.NewCollector()

	col.OnHTML("table#table > tbody >tr", func(ele *colly.HTMLElement) {
		ch := ele.DOM.Children()
		domain := ch.Eq(1).Text()
		address := ch.Eq(2).Children().Eq(0).Text()
		rtype := ch.Eq(3).Text()
		date := ch.Eq(4).Text()

		items = append(items, SameIPItem{domain, address, rtype, date})
	})

	url := "https://rapiddns.io/sameip/" + address + "?full=1"
	err := col.Visit(url)

	if err != nil {
		log.Fatalln(err)
	}

	return SameIpResult{items}
}

type SameIpResult struct {
	Items []SameIPItem `json:"items"`
}

type SameIPItem struct {
	Domain     string `json:"domain"`
	Address    string `json:"address"`
	RecordType string `json:"record_type"`
	Date       string `json:"date"`
}

type SameIPClient struct {
	Address *string
	Result  *SameIpResult
	S       Scraper[SameIpResult]
}

func (c *SameIPClient) Search() SameIpResult {
	if c.Address == nil {
		log.Fatalln("address is not specified.")
	}
	return c.S.Scrape(*c.Address)
}
