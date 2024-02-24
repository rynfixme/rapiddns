package main

type Scraper interface {
	Scrape() error
}
