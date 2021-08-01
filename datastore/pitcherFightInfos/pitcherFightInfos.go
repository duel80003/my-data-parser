package pitcherFightInfos

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

func (s Store) BatchUpsert(infos []*entities.PitcherFightInfos) {
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "player_id"}},
		UpdateAll: true,
	}).Create(&infos).Error
	if err != nil {
		logger.Errorf("BatchUpsert pitcherFightInfos error %d", err)
	} else {
		logger.Infoln("BatchUpsert pitcherFightInfos done")
	}

}
func (s Store) GetByPlayerId(playerId string) (infos []*entities.PitcherFightInfos, err error) {
	err = s.db.Select("id").Find(&infos, "player_id = ?", playerId).Error
	return infos, err
}
