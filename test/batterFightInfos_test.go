package test

import (
	"github.com/stretchr/testify/assert"
	"my-data-parser/datastore/batterFightInfos"
	"my-data-parser/driver"
	"my-data-parser/entities"
	"testing"
)

var batterFightStore = batterFightInfos.New(driver.DatabaseClient())

func createBatterInfo() *entities.BatterFightInfos {
	createUser()
	batterFightInfo := &entities.BatterFightInfos{
		PlayerId: PlayerID,
		Avg:      0.33,
		Acnt:     PlayerID,
	}
	batterFightStore.BatchUpsert([]*entities.BatterFightInfos{
		batterFightInfo,
	})

	info, _ := batterFightStore.GetByPlayerId(PlayerID)
	return info[0]
}

func TestBatterFightStore_BatchUpdate(t *testing.T) {
	info := createBatterInfo()
	ID := info.ID
	batterInfo := &entities.BatterFightInfos{
		PlayerId: PlayerID,
		Acnt:     PlayerID,
		Avg:      0.6,
	}

	batterInfo.SetID(ID)
	batterFightInfos := []*entities.BatterFightInfos{
		batterInfo,
	}

	batterFightStore.BatchUpsert(batterFightInfos)
	assert.Equal(t, batterInfo.ID, ID, "they should be equal")
	assert.Equal(t, batterInfo.Acnt, PlayerID, "they should be equal")
	assert.Equal(t, batterInfo.Avg, 0.6, "they should be equal")
}

func TestBatterFightStore_GetByPlayerId(t *testing.T) {
	playerId := "test1234"
	batterInfos, err := batterFightStore.GetByPlayerId(playerId)
	assert.Nil(t, err)
	for _, v := range batterInfos {
		assert.Equal(t, v.Avg, 0.0, "they should be equal")
		assert.Equal(t, v.Acnt, PlayerID, "they should be equal")
	}
}
