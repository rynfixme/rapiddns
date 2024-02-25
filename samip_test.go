package main

import "testing"

type SameIPTest struct {
	Name      string
	Args      SameIPTestArgs
	Exptected SameIPTestExpected
}

type SameIPTestArgs struct {
	Address string
}

type SameIPTestExpected struct {
	HasItems bool
}

var sameIPTests = []SameIPTest{
	SameIPTest{"Should get multiple items", SameIPTestArgs{"172.217.3.174"}, SameIPTestExpected{true}},
	SameIPTest{"Should get no items", SameIPTestArgs{"133.71.0.0/99"}, SameIPTestExpected{false}},
}

func TestSameIP(t *testing.T) {
	for _, tt := range sameIPTests {
		t.Run(tt.Name, func(t *testing.T) {
			sprov := SameIPScraperProvider[SameIpResult]{}
			c := SameIPClient{&tt.Args.Address, nil, &sprov}
			got := c.Search()

			if len(got.Items) > 0 != tt.Exptected.HasItems {
				t.Errorf("TestSameIP(), %v, HasItems error, %v", tt.Name, got.Items)
				return
			}
		})
	}
}
