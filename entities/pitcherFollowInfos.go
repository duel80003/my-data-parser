package entities

import "gorm.io/gorm"

type SId struct {
	Value string `gorm:"column:s_id"`
}
type PitcherFollowInfos struct {
	gorm.Model
	PlayerId                   string
	AssistCnt                  int
	BalkCnt                    int
	BasesONBallsCnt            int `gorm:"column:bases_on_balls_cnt"`
	CompleteGames              int
	EarnedRunCnt               int
	Era                        float64
	ErrorCnt                   int
	FieldNo                    string
	FightTeamAbbrName          string
	FightTeamCode              string
	FlyOut                     int
	GameDate                   string
	GameResult                 string
	GameSno                    int
	GroundOut                  int
	HitBYPitchCnt              int `gorm:"column:hit_by_pitch_cnt"`
	HittingCnt                 int
	HomeRunCnt                 int
	InningPitchedCnt           float64
	IntentionalBasesONBallsCnt int `gorm:"column:intentional_bases_on_balls_cnt"`
	JoinDoublePlayCnt          int
	JoinTripplePlayCnt         int
	KindCode                   string
	NoBaseBalled               int
	PitchCnt                   int
	PitchOutCnt                int
	PitcherAcnt                string
	PitcherName                string
	PlateAppearances           int
	PutoutCnt                  int
	RoleType                   string
	RunCnt                     int
	SaveFail                   int
	ShoutOut                   int
	StealCnt                   int
	StrikeCnt                  int
	StrikeOutCnt               int
	TeamNo                     string
	TotalTeamGames             int
	WildPitchCnt               int
	SId                        SId `gorm:"embedded"`
}

func (p *PitcherFollowInfos) SetID(id uint) {
	p.ID = id
}

func (b *PitcherFollowInfos) SetPlayerId(id string) {
	b.PlayerId = id
}
