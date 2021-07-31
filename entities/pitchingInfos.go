package entities

import "gorm.io/gorm"

type PitchingInfos struct {
	gorm.Model
	PlayerId                   string
	Year                       string
	TeamAbbrName               string
	Acnt                       string // player id
	TotalGames                 int    // Games Played
	PitchStarting              int    // S
	PitchCloser                int    // GF
	ShoutOut                   int    // SHO
	NoBaseBalled               int
	Wins                       int
	Loses                      int
	SaveOK                     int //SV
	SaveFail                   int //BSV
	ReliefPointCnt             int //HLD
	InningPitched              int //IP
	Whip                       float64
	Era                        float64
	PlateAppearances           int //TBF
	PitchCnt                   int //PC
	HittingCnt                 int
	HomeRunCnt                 int
	BasesONBallsCnt            int `gorm:"column:bases_on_balls_cnt"`
	IntentionalBasesONBallsCnt int `gorm:"column:intentional_bases_on_balls_cnt"` // IBB
	HitBYPitchCnt              int `gorm:"column:hit_by_pitch_cnt"`               // HBP
	StrikeOutCnt               int
	WildPitchCnt               int //WP
	BalkCnt                    int //BK
	RunCnt                     int //
	EarnedRunCnt               int // ER
	GroundOut                  int
	FlyOut                     int     //AO
	Goao                       float64 // GO/AO
	CompleteGames              int
}

func (p *PitchingInfos) SetID(id uint) {
	p.ID = id
}

func (p *PitchingInfos) SetPlayerId(id string) {
	p.PlayerId = id
}
