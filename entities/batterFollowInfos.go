package entities

import "gorm.io/gorm"

type BatterFollowInfos struct {
	gorm.Model
	PlayerId                   string
	Avg                        float64
	BasesONBallsCnt            int
	CaughtStealingCnt          int
	DoublePlayBatCnt           int
	ErrorCnt                   int
	FieldNo                    string
	FightTeamAbbrName          string
	FightTeamCode              string
	GameDate                   string
	GameSno                    int
	HitBYPitchCnt              int
	HitCnt                     int
	HitterAcnt                 string
	HitterName                 string
	HittingCnt                 int
	HomeRunCnt                 int
	IntentionalBasesONBallsCnt int `gorm:"column:intentional_bases_on_balls_cnt"`
	JoinDoublePlayCnt          int
	JoinTripplePlayCnt         int
	KindCode                   string
	Lobs                       int
	OneBaseHitCnt              int
	PassedBallCnt              int
	PlateAppearances           int
	PutoutCnt                  int
	RunBattedINCnt             int `gorm:"column:run_batted_in_cnt"`
	SacrificeFlyCnt            int
	SacrificeHitCnt            int
	ScoreCnt                   int
	StealBaseFailCnt           int
	StealBaseOKCnt             int `gorm:"column:steal_base_ok_cnt"`
	StrikeOutCnt               int
	TeamNo                     string
	ThreeBaseHitCnt            int
	TotalBases                 int
	TotalTeamGames             int
	TripplePlayBatCnt          int
	TwoBaseHitCnt              int
	SId                        SId `gorm:"embedded"`
}

func (b *BatterFollowInfos) SetID(id uint) {
	b.ID = id
}

func (b *BatterFollowInfos) SetPlayerId(id string) {
	b.PlayerId = id
}
