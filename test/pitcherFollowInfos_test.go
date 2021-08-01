package test

import (
	"github.com/stretchr/testify/assert"
	"my-data-parser/datastore/pitcherFollowInfos"
	"my-data-parser/driver"
	"my-data-parser/entities"
	"testing"
)

var pitcherFollowStore = pitcherFollowInfos.New(driver.DatabaseClient())

func createPitcherFollowData() *entities.PitcherFollowInfos {
	createUser()
	tmpPitcherFightInfo := &entities.PitcherFollowInfos{
		PlayerId: PlayerID,
		Era:      2.55,
	}
	pitcherFollowStore.BatchUpsert([]*entities.PitcherFollowInfos{
		tmpPitcherFightInfo,
	})
	info, _ := pitcherFollowStore.GetByPlayerId(PlayerID)
	return info[0]
}

func TestPitcherFollowStore_BatchUpdate(t *testing.T) {
	createPitcherFollowData()
	Sid := "test5678"
	pitcherFollowInfo := &entities.PitcherFollowInfos{
		PlayerId: PlayerID,
		Era:      3.13,
		SId:      entities.SId{Value: Sid},
	}
	ID := pitcherFollowInfo.ID
	pitcherFollowInfo.SetID(ID)
	pitcherFightInfos := []*entities.PitcherFollowInfos{
		pitcherFollowInfo,
	}

	pitcherFollowStore.BatchUpsert(pitcherFightInfos)
	assert.Equal(t, pitcherFollowInfo.PlayerId, PlayerID, "they should be equal")
	assert.Equal(t, pitcherFollowInfo.Era, 3.13, "they should be equal")
	assert.Equal(t, pitcherFollowInfo.SId.Value, Sid, "they should be equal")
}

func TestStore_GetByPlayerId(t *testing.T) {
	createPitcherFollowData()
	pitchFightInfos, err := batterFollowStore.GetByPlayerId(PlayerID)
	assert.Nil(t, err)
	for _, v := range pitchFightInfos {
		t.Logf("%+v", v)
		assert.Equal(t, v.SId.Value, PlayerID, "they should be equal")
	}
}
