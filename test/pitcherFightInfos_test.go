package test

import (
	"github.com/stretchr/testify/assert"
	"my-data-parser/datastore/pitcherFightInfos"
	"my-data-parser/driver"
	"my-data-parser/entities"
	"testing"
)

var pitcherFightStore = pitcherFightInfos.New(driver.DatabaseClient())

func createPitcherFightData() *entities.PitcherFightInfos {
	createUser()
	tmpPitcherFightInfo := &entities.PitcherFightInfos{
		PlayerId: PlayerID,
		Era:      2.55,
	}
	pitcherFightStore.BatchUpsert([]*entities.PitcherFightInfos{
		tmpPitcherFightInfo,
	})
	info, _ := pitcherFightStore.GetByPlayerId(PlayerID)
	return info[0]
}

func TestPitcherFightStore_BatchUpdate(t *testing.T) {
	info := createPitcherFightData()
	ID := info.ID
	pitcherFightInfo := &entities.PitcherFightInfos{
		PlayerId: PlayerID,
		Era:      3.13,
	}
	pitcherFightInfo.SetID(ID)
	pitcherFightInfos := []*entities.PitcherFightInfos{
		pitcherFightInfo,
	}
	pitcherFightStore.BatchUpsert(pitcherFightInfos)
	assert.Equal(t, pitcherFightInfo.ID, uint(ID), "they should be equal")
	assert.Equal(t, pitcherFightInfo.PlayerId, PlayerID, "they should be equal")
	assert.Equal(t, pitcherFightInfo.Era, 3.13, "they should be equal")
}

func TestPitcherFightStore_GetByPlayerId(t *testing.T) {
	pitchFightInfos, err := pitcherFightStore.GetByPlayerId(PlayerID)
	assert.Nil(t, err)
	for _, v := range pitchFightInfos {
		t.Logf("%+v", v)
		assert.NotNil(t, v)
	}
}
