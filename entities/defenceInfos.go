package entities

import "gorm.io/gorm"

type DefenceInfos struct {
	gorm.Model
	PlayerId           string
	Year               string
	TeamNo             string
	Fpct               float64 // FP
	Acnt               string  // player id
	DefendStationName  string  // position
	TotalGames         int
	DefendCnt          int // TC
	PutoutCnt          int // PO
	AssistCnt          int // A
	ErrorCnt           int // E
	JoinDoublePlayCnt  int // DP
	JoinTripplePlayCnt int // TP
}

func (d *DefenceInfos) SetID(id uint) {
	d.ID = id
}

func (d *DefenceInfos) SetPlayerId(id string) {
	d.PlayerId = id
}
