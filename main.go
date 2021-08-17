package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
	"github.com/urmandor/web-scrapper/api"
)

func main() {
	c := colly.NewCollector()

	fixtures := api.GetAllFixtures(c.Clone())
	data, err := json.MarshalIndent(fixtures, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ioutil.WriteFile("fixtures.json", data, 0644)

	results := api.GetAllResults(c.Clone())
	data, err = json.MarshalIndent(results, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ioutil.WriteFile("results.json", data, 0644)

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("visiting " + r.URL.String() + "\n")
	// })
	// c.OnHTML(".match-scorecard-table", func(h *colly.HTMLElement) {
	// 	fmt.Println(h.Name)
	// })
	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("All done...\n")
	// })
	// c.Visit(constants.GenerateURL(constants.Series, constants.Match))

}
