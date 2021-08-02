package standingInfos

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

func (s Store) BatchUpsert(infos []*entities.StandingInfos) {
	err := s.db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&infos).Error
	if err != nil {
		logger.Errorf("BatchUpsert standinginfos error %s", err)
	} else {
		logger.Infoln("BatchUpsert standingInfos done")
	}
}

func (s Store) GetAll() (infos []*entities.StandingInfos) {
	s.db.Debug().Find(&infos)
	return infos
}
