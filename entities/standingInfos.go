package entities

import "gorm.io/gorm"

type StandingInfos struct {
	gorm.Model
	TeamName          string
	SortNumber        int
	Games             int
	WinAndLose        string
	GamesBehind       string
	WinningPercentage float64
	GamesEliminate    int
	Last10Games       string `gorm:"column:last_10_games"`
	RecordAtHome      string
	RecordAtAway      string
	CurrentStreak     string
	Brothers          string
	ULions            string
	Monkeys           string
	Guardians         string
	Dragons           string
}

func (s *StandingInfos) SetId(id uint) {
	s.ID = id
}
