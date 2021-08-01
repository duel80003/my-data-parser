package defenceInfos

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"my-data-parser/entities"
	"my-data-parser/utils"
)

var logger = utils.LoggerInstance()

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) Store {
	return Store{db: db}
}

func (s Store) BatchUpsert(defenceInfos []*entities.DefenceInfos) {
	if defenceInfos == nil || len(defenceInfos) == 0 {
		return
	}
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "player_id"}},
		UpdateAll: true,
	}).Create(&defenceInfos).Error
	if err != nil {
		logger.Errorf("BatchUpsert defenceInfos error %d", err)
	} else {
		logger.Infoln("BatchUpsert defenceInfos done")
	}
}

func (s Store) GetByPlayerId(playerId string) (infos []*entities.DefenceInfos, err error) {
	err = s.db.Select("id", "year", "defend_station_name").Find(&infos, "player_id = ?", playerId).Error
	return infos, err
}
