package test

import (
	"github.com/stretchr/testify/assert"
	"my-data-parser/datastore/pitchingInfos"
	"my-data-parser/driver"
	"my-data-parser/entities"
	"testing"
)

var pitchingStore = pitchingInfos.New(driver.DatabaseClient())

func createPitchInfo() *entities.PitchingInfos {
	createUser()
	pitchInfo := &entities.PitchingInfos{
		PlayerId: PlayerID,
		Era:      3.22,
		Year:     "2021",
		Acnt:     PlayerID,
	}
	pitchingStore.BatchUpsert([]*entities.PitchingInfos{
		pitchInfo,
	})
	info, _ := pitchingStore.GetByPlayerId(PlayerID)
	return info[0]
}

func TestPitchingStoreStore_BatchUpdate(t *testing.T) {
	info := createPitchInfo()
	pitchingInfo := &entities.PitchingInfos{
		PlayerId: info.PlayerId,
		Era:      3.13,
		Year:     "2021",
		Acnt:     info.PlayerId,
	}
	ID := info.ID
	pitchingInfo.SetID(ID)
	pitcherInfos := []*entities.PitchingInfos{
		pitchingInfo,
	}

	pitchingStore.BatchUpsert(pitcherInfos)
	assert.Equal(t, pitchingInfo.ID, uint(ID), "they should be equal")
	assert.Equal(t, pitchingInfo.PlayerId, PlayerID, "they should be equal")
	assert.Equal(t, pitchingInfo.Era, 3.13, "they should be equal")
	assert.Equal(t, pitchingInfo.Year, "2021", "they should be equal")
	assert.Equal(t, pitchingInfo.Acnt, pitchingStore, "they should be equal")
}

func TestPitchingStoreStore_GetByPlayerId(t *testing.T) {
	info := createPitchInfo()
	ID := info.ID
	pitchingInfos, err := pitchingStore.GetByPlayerId(PlayerID)
	assert.Nil(t, err)
	for _, v := range pitchingInfos {
		t.Logf("%+v", v)
		assert.Equal(t, v.ID, uint(ID), "they should be equal")
		assert.Equal(t, v.Year, "2021", "they should be equal")
	}
}
