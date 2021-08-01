package datastore

import (
	"my-data-parser/entities"
)

type IDBPlayers interface {
	GetPlayers() (players []*entities.Player, err error)
	UpsertWithoutUpdate(player *entities.Player)
	Upsert(player *entities.Player)
}

type IStandingInfos interface {
	BatchUpsert(infos []*entities.StandingInfos)
}

type IBattingInfos interface {
	BatchUpsert(infos []*entities.BattingInfos)
	GetByPlayerId(playerId string) (infos []*entities.BattingInfos, err error)
}

//
type IDefenceInfos interface {
	BatchUpsert(infos []*entities.DefenceInfos)
	GetByPlayerId(playerId string) (infos []*entities.DefenceInfos, err error)
}

type IBatterFollowInfos interface {
	BatchUpsert(infos []*entities.BatterFollowInfos)
	GetByPlayerId(playerId string) (infos []*entities.BatterFollowInfos, err error)
}

type IPitchingInfos interface {
	BatchUpsert(infos []*entities.PitchingInfos)
	GetByPlayerId(playerId string) (infos []*entities.PitchingInfos, err error)
}

type IPitcherFightInfos interface {
	BatchUpsert(infos []*entities.PitcherFightInfos)
	GetByPlayerId(playerId string) (infos []*entities.PitcherFightInfos, err error)
}

type IPitcherFollowInfos interface {
	BatchUpsert(infos []*entities.PitcherFollowInfos)
	GetByPlayerId(playerId string) (infos []*entities.PitcherFollowInfos, err error)
}

type IBatterFightInfos interface {
	BatchUpsert(infos []*entities.BatterFightInfos)
	GetByPlayerId(playerId string) (infos []*entities.BatterFightInfos, err error)
}
