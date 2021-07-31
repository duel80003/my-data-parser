package entities

import "gorm.io/gorm"

type PitcherFightInfos struct {
	gorm.Model
	PlayerId                   string
	Acnt                       string
	BalkCnt                    int
	BasesONBallsCnt            int `gorm:"column:bases_on_balls_cnt"`
	CompleteGames              int
	EarnedRunCnt               int // ER
	Enabled                    bool
	Era                        float64
	FightTeamCode              string
	FightTeamName              string
	HitBYPitchCnt              int `gorm:"column:hit_by_pitch_cnt"`
	HittingCnt                 int
	HomeRunCnt                 int
	InningPitched              int
	InningPitchedCnt           int
	InningPitchedDiv3Cnt       int `gorm:"column:inning_pitched_div_3_cnt"`
	IntentionalBasesONBallsCnt int `gorm:"column:intentional_bases_on_balls_cnt"`
	Loses                      int
	PitchCloser                int
	PitchCnt                   int
	PitchStarting              int
	PlateAppearances           int
	ReliefPointCnt             int
	RunCnt                     int
	SaveFail                   int
	SaveOK                     int
	ShoutOut                   int
	StrikeOutCnt               int
	TeamNo                     string
	TotalGames                 int
	Whip                       float64
	WildPitchCnt               int
	Wins                       int
}

func (p *PitcherFightInfos) SetID(id uint) {
	p.ID = id
}

func (p *PitcherFightInfos) SetPlayerId(id string) {
	p.PlayerId = id
}
