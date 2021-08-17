package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/urmandor/web-scrapper/constants"
)

type Fixture struct {
	MatchId      int
	HomeTeamName string
	AwayTeamName string
	Date         string
	MatchNumber  int
	Venue        string
	URL          string
}

func ExtractFixture(h *colly.HTMLElement) Fixture {
	description := h.ChildText(".description")

	split := strings.Split(description, ",")
	re := regexp.MustCompile(`^[0-9]+`)
	match, _ := strconv.Atoi(re.FindString(split[0]))

	venue := split[1]
	date := split[2]

	homeTeam := ""
	awayTeam := ""
	h.ForEach(".teams .team .name-detail p.name", func(i int, e *colly.HTMLElement) {
		if i == 0 {
			homeTeam = e.Text
		} else {
			awayTeam = e.Text
		}
	})

	query := h.DOM
	url, _ := query.Find(".match-info-link-FIXTURES").Attr("href")
	re = regexp.MustCompile(`[0-9]+/[a-z\-]+$`)
	matchId, _ := strconv.Atoi(strings.Split(re.FindString(url), "/")[0])
	fixture := Fixture{
		MatchId:      matchId,
		HomeTeamName: homeTeam,
		AwayTeamName: awayTeam,
		Date:         date,
		URL:          url,
		MatchNumber:  match,
		Venue:        venue,
	}

	return fixture

}

func GetAllFixtures(c *colly.Collector) []Fixture {
	var wg sync.WaitGroup
	fixtures := []Fixture{}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting " + r.URL.String() + "\n")
		wg.Add(1)
	})
	c.OnHTML(".match-info.match-info-FIXTURES", func(h *colly.HTMLElement) {
		fixture := ExtractFixture(h)
		fixtures = append(fixtures, fixture)
	})
	c.OnScraped(func(r *colly.Response) {
		defer wg.Done()
		fmt.Println("All done...\n")
	})
	c.Visit(constants.GenerateURL(constants.Series, constants.Fixtures))

	wg.Wait()
	return fixtures
}
