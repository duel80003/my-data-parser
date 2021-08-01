package batterFightInfos

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

func (s Store) BatchUpsert(battingInfos []*entities.BatterFightInfos) {
	err := s.db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&battingInfos).Error
	if err != nil {
		logger.Errorf("BatchUpsert batterFightInfos error %s", err)
	} else {
		logger.Infoln("BatchUpsert batterFightInfos done")
	}
}

func (s Store) GetByPlayerId(playerId string) (infos []*entities.BatterFightInfos, err error) {
	err = s.db.Select("id").Find(&infos, "player_id = ?", playerId).Error
	return infos, err
}
