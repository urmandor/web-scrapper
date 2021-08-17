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

type Result struct {
	MatchId     int
	WinningTeam Score
	LosingTeam  Score
	Result      string
	Date        string
	MatchNumber int
	Venue       string
	URL         string
}

type Score struct {
	Name    string
	Runs    int
	Wickets int
	Balls   int
}

func extractScore(e *colly.HTMLElement) Score {
	var runs, wickets int
	name := e.ChildText(".name-detail p.name")
	score := strings.Split(e.ChildText(".score-detail .score"), "/")

	if len(score) == 2 {
		runs, _ = strconv.Atoi(score[0])
		wickets, _ = strconv.Atoi(score[1])
	}

	return Score{
		Name: name, Runs: runs, Wickets: wickets,
	}
}

func ExtractResult(h *colly.HTMLElement) Result {
	description := h.ChildText(".description")
	status := h.ChildText(".status-text span")
	split := strings.Split(description, ",")
	re := regexp.MustCompile(`^[0-9]+`)
	match, _ := strconv.Atoi(re.FindString(split[0]))

	venue := split[1]
	date := split[2]

	winningTeamScore := Score{}
	losingTeamScore := Score{}
	h.ForEach(".teams .team", func(i int, e *colly.HTMLElement) {
		if i == 0 {
			winningTeamScore = extractScore(e)
		} else {
			losingTeamScore = extractScore(e)
		}
	})

	query := h.DOM
	url, _ := query.Find(".match-info-link-FIXTURES").Attr("href")
	re = regexp.MustCompile(`[0-9]+/[a-z\-]+$`)
	matchId, _ := strconv.Atoi(strings.Split(re.FindString(url), "/")[0])
	result := Result{
		MatchId:     matchId,
		Result:      status,
		WinningTeam: winningTeamScore,
		LosingTeam:  losingTeamScore,
		Date:        date,
		URL:         url,
		MatchNumber: match,
		Venue:       venue,
	}

	return result

}

func GetAllResults(c *colly.Collector) []Result {
	var wg sync.WaitGroup
	results := []Result{}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting " + r.URL.String() + "\n")
		wg.Add(1)
	})
	c.OnHTML(".match-info.match-info-FIXTURES", func(h *colly.HTMLElement) {
		result := ExtractResult(h)
		results = append(results, result)
	})
	c.OnScraped(func(r *colly.Response) {
		defer wg.Done()
		fmt.Println("All done...\n")
	})
	c.OnError(func(r *colly.Response, e error) {
		defer wg.Done()
		fmt.Println(e)
	})
	c.Visit(constants.GenerateURL(constants.Series, constants.Results))

	wg.Wait()
	return results
}
