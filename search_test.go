package main

import "testing"

type SearchTest struct {
	Name      string
	Args      SearchTestArgs
	Exptected SearchTestExpected
}

type SearchTestArgs struct {
	Word string
}

type SearchTestExpected struct {
	HasItems bool
}

var SearchTests = []SearchTest{
	SearchTest{"Should get multiple items", SearchTestArgs{"redbull"}, SearchTestExpected{true}},
	SearchTest{"Should get no items", SearchTestArgs{"rynfixme"}, SearchTestExpected{false}},
}

func TestSearch(t *testing.T) {
	for _, tt := range SearchTests {
		t.Run(tt.Name, func(t *testing.T) {
			sprov := SearchScraperProvider[SearchResult]{}
			c := SearchClient{&tt.Args.Word, nil, &sprov}
			got := c.Search()

			if len(got.Items) > 0 != tt.Exptected.HasItems {
				t.Errorf("TestSearch(), %v, HasItems error, %v", tt.Name, got.Items)
				return
			}
		})
	}
}
