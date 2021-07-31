package entities

import "gorm.io/gorm"

type StandingInfos struct {
	gorm.Model
	SortNumber        int
	Games             int
	WinAndLose        string
	GamesBehind       string
	WinningPercentage float64
	GamesEliminate    int
	Last10Games       string
	RecordAtHome      string
	RecordAtAway      string
	CurrentStace      string
	Brothers          string
	ULions            string
	Monkeys           string
	Guardians         string
	Dragons           string
}
