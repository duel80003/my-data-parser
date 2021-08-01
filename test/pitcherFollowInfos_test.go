package test

import (
	"my-data-parser/datastore/pitcherFollowInfos"
	"my-data-parser/driver"
)

var pitcherFollowStore = pitcherFollowInfos.New(driver.DatabaseClient())

//const ID = 7149
//const playerId = "test5678"
//
//func TestStore_BatchUpdate(t *testing.T) {
//	pitcherFollowInfo := &entities.PitcherFollowInfos{
//		PlayerId: playerId,
//		Era:      3.13,
//		SId:      entities.SId{Value: "test5678"},
//	}
//	pitcherFollowInfo.SetID(ID)
//	pitcherFightInfos := []*entities.PitcherFollowInfos{
//		pitcherFollowInfo,
//	}
//
//	store.BatchUpsert(pitcherFightInfos)
//	assert.Equal(t, pitcherFollowInfo.ID, uint(ID), "they should be equal")
//	assert.Equal(t, pitcherFollowInfo.PlayerId, playerId, "they should be equal")
//	assert.Equal(t, pitcherFollowInfo.Era, 3.13, "they should be equal")
//	assert.Equal(t, pitcherFollowInfo.SId.Value, playerId, "they should be equal")
//}
//
//func TestStore_GetByPlayerId(t *testing.T) {
//	pitchFightInfos, err := store.GetByPlayerId(playerId)
//	assert.Nil(t, err)
//	for _, v := range pitchFightInfos {
//		t.Logf("%+v", v)
//		assert.Equal(t, v.ID, uint(ID), "they should be equal")
//		assert.Equal(t, v.SId.Value, playerId, "they should be equal")
//
//	}
//}
