package constants

import (
	"strings"
)

const (
	BaseURL  = "https://www.espncricinfo.com/series"
	Series   = "kashmir-premier-league-2021-1272105"
	Fixtures = "match-schedule-fixtures"
	Results  = "match-results"
	Match    = "rawalakot-hawks-vs-muzzaffarabad-tigers-7th-match-1272121"
)

func GenerateURL(names ...string) string {
	complete := append([]string{BaseURL}, names...)
	return strings.Join(complete, "/")
}
