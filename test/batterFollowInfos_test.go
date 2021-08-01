package test

import (
	"github.com/stretchr/testify/assert"
	"my-data-parser/datastore/batterFollowInfos"
	"my-data-parser/driver"
	"my-data-parser/entities"
	"testing"
)

var batterFollowStore = batterFollowInfos.New(driver.DatabaseClient())

func createBatterFollow() *entities.BatterFollowInfos {
	createUser()
	batterInfos := &entities.BatterFollowInfos{
		PlayerId: PlayerID,
		Avg:      0.33,
		SId:      entities.SId{Value: "test1234"},
	}
	batterFollowStore.BatchUpsert([]*entities.BatterFollowInfos{
		batterInfos,
	})
	info, _ := batterFollowStore.GetByPlayerId(PlayerID)
	return info[0]

}

func TestBatterFollowStore_BatchUpdate(t *testing.T) {
	info := createBatterFollow()
	batterInfo := &entities.BatterFollowInfos{
		PlayerId: info.PlayerId,
		SId:      entities.SId{Value: "test1234"},
		Avg:      0.6,
	}

	batterInfo.SetID(info.ID)
	batterFollowInfos := []*entities.BatterFollowInfos{
		batterInfo,
	}

	batterFollowStore.BatchUpsert(batterFollowInfos)
	assert.Equal(t, batterInfo.ID, info.ID, "they should be equal")
	assert.Equal(t, batterInfo.SId.Value, "test1234", "they should be equal")
	assert.Equal(t, batterInfo.Avg, 0.6, "they should be equal")
}

func TestBatterFollowStore_GetByPlayerId(t *testing.T) {
	playerId := "test1234"
	batterInfos, err := batterFollowStore.GetByPlayerId(playerId)
	assert.Nil(t, err)
	for _, v := range batterInfos {
		t.Logf("%+v", v)
		assert.Equal(t, v.ID, uint(2408), "they should be equal")
		assert.Equal(t, v.SId.Value, "test1234", "they should be equal")
		assert.Equal(t, v.Avg, 0.0, "they should be equal")
	}
}
