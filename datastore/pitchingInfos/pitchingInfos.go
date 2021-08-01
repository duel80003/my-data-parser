package pitchingInfos

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

func (s Store) BatchUpsert(infos []*entities.PitchingInfos) {
	err := s.db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&infos).Error
	if err != nil {
		logger.Errorf("BatchUpsert pitchingInfos error %s", err)
	} else {
		logger.Infoln("BatchUpsert pitchingInfos done")
	}
}

func (s Store) GetByPlayerId(playerId string) (infos []*entities.PitchingInfos, err error) {
	err = s.db.Select("id", "year").Find(&infos, "player_id = ?", playerId).Error
	return infos, err
}
