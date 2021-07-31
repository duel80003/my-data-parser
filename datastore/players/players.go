package players

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

func (s Store) GetPlayers() (players []*entities.Player, err error) {
	err = s.db.Find(&players).Error
	return players, err
}

func (s Store) Upsert(player *entities.Player) {
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&player).Error
	if err != nil {
		logger.Errorf("BatchUpsertWithoutUpdate error %s", err)
	} else {
		logger.Infoln("BatchUpsertWithoutUpdate done")
	}
}
