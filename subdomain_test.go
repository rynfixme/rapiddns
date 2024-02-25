package main

import "testing"

type SubdomainTest struct {
	Name      string
	Args      SubdomainTestArgs
	Exptected SubdomainTestExpected
}

type SubdomainTestArgs struct {
	Address string
}

type SubdomainTestExpected struct {
	HasItems bool
}

var SubdomainTests = []SubdomainTest{
	SubdomainTest{"Should get multiple items", SubdomainTestArgs{"tesla.com"}, SubdomainTestExpected{true}},
	SubdomainTest{"Should get no items", SubdomainTestArgs{"rynfixme.com"}, SubdomainTestExpected{false}},
}

func TestSubdomain(t *testing.T) {
	for _, tt := range SubdomainTests {
		t.Run(tt.Name, func(t *testing.T) {
			sprov := SubdomainScraperProvider[SubdomainResult]{}
			c := SubdomainClient{&tt.Args.Address, nil, &sprov}
			got := c.Search()

			if len(got.Items) > 0 != tt.Exptected.HasItems {
				t.Errorf("TestSubdomain(), %v, HasItems error, %v", tt.Name, got.Items)
				return
			}
		})
	}
}
