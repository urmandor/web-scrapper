package api

type DismissalType int

const (
	Bowled DismissalType = iota
	LBW
	Caught
	Runout
	Stumped
	HitWicket
)

type BattingCard struct {
	Name        string
	Dismissal   DismissalType
	DismissedBy string
	Runs        int
	Balls       int
	Minutes     int
	Fours       int
	Sixes       int
}

// func ExtractBattingScore(h *colly.HTMLElement) Result {

// }
