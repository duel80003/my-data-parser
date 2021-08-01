package test

import (
	"github.com/stretchr/testify/assert"
	"my-data-parser/datastore/battingInfos"
	"my-data-parser/driver"
	"my-data-parser/entities"
	"testing"
)

var battingStore = battingInfos.New(driver.DatabaseClient())

func createBattingInfo() *entities.BattingInfos {
	createUser()
	battingInfo := &entities.BattingInfos{
		PlayerId: PlayerID,
		Avg:      0.33,
	}
	battingStore.BatchUpsert([]*entities.BattingInfos{
		battingInfo,
	})

	info, _ := battingStore.GetByPlayerId(PlayerID)
	return info[0]
}

func TestBattingStore_BatchUpdate(t *testing.T) {
	info := createBattingInfo()
	batterInfo := &entities.BattingInfos{
		PlayerId: PlayerID,
		Avg:      0.6,
	}
	batterInfo.SetID(info.ID)
	batterInfos := []*entities.BattingInfos{
		batterInfo,
	}

	battingStore.BatchUpsert(batterInfos)
	assert.Equal(t, batterInfo.ID, info.ID, "they should be equal")
	assert.Equal(t, batterInfo.Avg, 0.6, "they should be equal")
}

func TestBattingStore_GetByPlayerId(t *testing.T) {
	batterInfos, err := battingStore.GetByPlayerId(PlayerID)
	assert.Nil(t, err)
	for _, v := range batterInfos {
		t.Logf("%+v", v)
		assert.Equal(t, v.Avg, 0.0, "they should be equal")
		assert.Equal(t, v.Year, "", "they should be equal")
	}
}
