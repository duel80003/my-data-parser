package entities

import "gorm.io/gorm"

type BattingInfos struct {
	gorm.Model
	PlayerId                   string
	TeamAbbrName               string
	Obp                        float64 // OBP
	Slg                        float64 // SLG
	Avg                        float64 // Batting Average
	Goao                       float64 // GO/AO
	Ops                        float64
	Acnt                       string // player id
	Year                       string
	TotalGames                 int // Games Played
	PlateAppearances           int // PA
	HitCnt                     int // AB
	RunBattedINCnt             int `gorm:"column:run_batted_in_cnt"` // RBI
	ScoreCnt                   int // Runs
	HittingCnt                 int // H
	OneBaseHitCnt              int // Single
	TwoBaseHitCnt              int // Double
	ThreeBaseHitCnt            int // Triple
	HomeRunCnt                 int
	TotalBases                 int // TB
	StrikeOutCnt               int
	StealBaseOKCnt             int // SB
	DoublePlayBatCnt           int // DP
	SacrificeHitCnt            int // SH
	SacrificeFlyCnt            int // SF
	BasesONBallsCnt            int // Walk
	IntentionalBasesONBallsCnt int // IBB
	HitBYPitchCnt              int // HBP
	StealBaseFailCnt           int // CS
	GroundOut                  int
	FlyOut                     int //AO
	SB                         int // SB%
}

func (b *BattingInfos) SetID(id uint) {
	b.ID = id
}

func (b *BattingInfos) SetPlayerId(id string) {
	b.PlayerId = id
}
