package main

import (
	"encoding/json"
	"fmt"
	"os"

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
		if searchWord != nil {
			prov := SearchScraperProvider[SearchItem]{}
			search := SearchClient{*searchWord, SearchResult{[]SearchItem{}}, &prov}
			search.Search()
			bytes, err := json.Marshal(search.Result)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(bytes))
			return
		}

		fmt.Println(app.Help)

	case sameIP.FullCommand():
		if sameIPAdress != nil {
			prov := SameIPScraperProvider[SameIPItem]{}
			ip := SameIPClient{*sameIPAdress, SameIpResult{[]SameIPItem{}}, &prov}
			ip.GetSameIP()
			bytes, err := json.Marshal(ip.Result)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(bytes))
			return
		}

		fmt.Println(app.Help)

	case subdomain.FullCommand():
		if subdomainDomain != nil {
			prov := SubdomainScraperProvider[SubdomainItem]{}
			subdomain := SubdomainClient{*subdomainDomain, SubdomainResult{[]SubdomainItem{}}, &prov}
			subdomain.GetSubdomain()
			bytes, err := json.Marshal(subdomain.Result)

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
