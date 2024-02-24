package main

import (
	"log"

	"github.com/gocolly/colly"
)

type SameIPScraperProvider[T SameIPItem] struct{}

func (s *SameIPScraperProvider[SameIPItem]) Scrape(address string) []SameIPItem {
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

	return items
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
	Address string
	Result  SameIpResult
	S       Scraper[SameIPItem]
}

func (sc *SameIPClient) getSameIP() error {
	domains := sc.S.Scrape(sc.Address)
	sc.Result.Items = domains
	return nil
}
