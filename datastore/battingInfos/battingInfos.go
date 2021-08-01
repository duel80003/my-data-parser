package battingInfos

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

func (s Store) BatchUpsert(battingInfos []*entities.BattingInfos) {
	err := s.db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&battingInfos).Error
	if err != nil {
		logger.Errorf("BatchUpsert battingInfos error %s", err)
	} else {
		logger.Infoln("BatchUpsert battingInfos done")
	}
}

func (s Store) GetByPlayerId(playerId string) (infos []*entities.BattingInfos, err error) {
	err = s.db.Select("id", "year").Find(&infos, "player_id = ?", playerId).Error
	return infos, err
}

//
//func (s Store) BatchUpsert(battingInfos []*entities.Info) {
//	err := s.db.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "playerId"}},
//		UpdateAll: true,
//		//DoUpdates: clause.AssignmentColumns([]string{"number", "team"}),
//	}).Create(&battingInfos).Error
//	if err != nil {
//		logger.Errorf("BatchUpsert battingInfos error %s", err)
//	} else {
//		logger.Infoln("BatchUpsert battingInfos done")
//	}
//}
