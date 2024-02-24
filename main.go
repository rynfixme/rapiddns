package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.New("rapiddns", "RapidDNS Scraping tool.")

	search     = app.Command("search", "Search by word.")
	searchWord = search.Arg("word", "Word to find").Required().String()
	// searchOutFile = search.Arg("out", "File to output")

	sameIP       = app.Command("sameip", "Search by IP address, CIDR.")
	sameIPAdress = sameIP.Arg("address", "Address or CIDR to find").Required().String()
	// sameIPOutFile = search.Arg("out", "File to output")

	subdomain       = app.Command("subdomain", "Search by Subdomain.")
	subdomainDomain = subdomain.Arg("domain", "domain to find").Required().String()
	// subdomainOutfile = search.Arg("out", "File to output")

)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case search.FullCommand():
		sw := searchWord
		word := strings.Clone(*sw)
		prov := SearchScraperProvider[SearchItem]{}
		search := SearchClient{word, SearchResult{[]SearchItem{}}, &prov}
		search.Search()
		bytes, err := json.Marshal(search.Result)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(bytes))

	case sameIP.FullCommand():
		sa := sameIPAdress
		address := strings.Clone(*sa)
		prov := SameIPScraperProvider[SameIPItem]{}
		ip := SameIPClient{address, SameIpResult{[]SameIPItem{}}, &prov}
		ip.GetSameIP()
		bytes, err := json.Marshal(ip.Result)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(bytes))

	case subdomain.FullCommand():
		d := subdomainDomain
		domain := strings.Clone(*d)
		prov := SubdomainScraperProvider[SubdomainItem]{}
		subdomain := SubdomainClient{domain, SubdomainResult{[]SubdomainItem{}}, &prov}
		subdomain.GetSubdomain()
		bytes, err := json.Marshal(subdomain.Result)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(bytes))

	default:
		fmt.Println(app.Help)
	}

}
