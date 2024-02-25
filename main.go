package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.New("rapiddns", "RapidDNS Scraping tool.")

	search     = app.Command("search", "Search by word.")
	searchWord = search.Arg("word", "Word to find").Required().String()

	sameIP       = app.Command("sameip", "Search by IP address, CIDR.")
	sameIPAdress = sameIP.Arg("address", "Address or CIDR to find").Required().String()

	subdomain       = app.Command("subdomain", "Search by Subdomain.")
	subdomainDomain = subdomain.Arg("domain", "domain to find").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case search.FullCommand():
		if searchWord != nil {
			prov := SearchScraperProvider[SearchResult]{}
			c := SearchClient{searchWord, nil, &prov}

			rerult := c.Search()
			c.Result = &rerult
			if c.Result == nil {
				log.Fatalln("Search has not completed")
			}

			bytes, err := json.Marshal(c.Result)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(bytes))
			return
		}

		fmt.Println(app.Help)

	case sameIP.FullCommand():
		if sameIPAdress != nil {
			prov := SameIPScraperProvider[SameIpResult]{}
			c := SameIPClient{sameIPAdress, nil, &prov}

			rerult := c.Search()
			c.Result = &rerult
			if c.Result == nil {
				log.Fatalln("Search same IP has not completed")
			}

			bytes, err := json.Marshal(c.Result)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(bytes))
			return
		}

		fmt.Println(app.Help)

	case subdomain.FullCommand():
		if subdomainDomain != nil {
			prov := SubdomainScraperProvider[SubdomainResult]{}
			c := SubdomainClient{subdomainDomain, nil, &prov}

			rerult := c.Search()
			c.Result = &rerult
			if c.Result == nil {
				log.Fatalln("Search same IP has not completed")
			}

			bytes, err := json.Marshal(c.Result)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(bytes))
			return
		}

		fmt.Println(app.Help)

	default:
		fmt.Println(app.Help)
	}

}
