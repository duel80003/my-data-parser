package entities

import "gorm.io/gorm"

type BatterFightInfos struct {
	gorm.Model
	PlayerId                   string
	Acnt                       string
	Avg                        float64
	BasesONBallsCnt            int `gorm:"column:bases_on_balls_cnt"`
	DoublePlayBatCnt           int
	Enabled                    bool
	FightTeamCode              string
	FightTeamName              string
	HitBYPitchCnt              int `gorm:"column:hit_by_pitch_cnt"`
	HitCnt                     int
	HittingCnt                 int
	HomeRunCnt                 int
	IntentionalBasesONBallsCnt int `gorm:"column:intentional_bases_on_balls_cnt"`
	Obp                        float64
	OneBaseHitCnt              int
	Ops                        float64
	PlateAppearances           int
	RunBattedINCnt             int `gorm:"column:run_batted_in_cnt"`
	SB                         int
	SacrificeFlyCnt            int
	SacrificeHitCnt            int
	ScoreCnt                   int
	Slg                        float64
	StealBaseFailCnt           int
	StealBaseOKCnt             int `grom:"column:steal_base_ok_cnt"`
	StrikeOutCnt               int
	TA                         float64
	TeamNo                     string
	ThreeBaseHitCnt            int
	TotalBases                 int
	TotalGames                 int
	TwoBaseHitCnt              int
}

func (b *BatterFightInfos) SetID(id uint) {
	b.ID = id
}
func (b *BatterFightInfos) SetPlayerId(id string) {
	b.PlayerId = id
}
