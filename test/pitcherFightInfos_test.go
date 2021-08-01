package test

import (
	"my-data-parser/datastore/pitcherFightInfos"
	"my-data-parser/driver"
)

var pitcherFightStore = pitcherFightInfos.New(driver.DatabaseClient())

//const ID = 12963
//const playerId = "test5678"
//
//func TestStore_BatchUpdate(t *testing.T) {
//	pitcherFightInfo := &entities.PitcherFightInfos{
//		PlayerId: playerId,
//		Era:      3.13,
//	}
//	pitcherFightInfo.SetID(ID)
//	pitcherFightInfos := []*entities.PitcherFightInfos{
//		pitcherFightInfo,
//	}
//	pitcherFightStore.BatchUpsert(pitcherFightInfos)
//	assert.Equal(t, pitcherFightInfo.ID, uint(ID), "they should be equal")
//	assert.Equal(t, pitcherFightInfo.PlayerId, playerId, "they should be equal")
//	assert.Equal(t, pitcherFightInfo.Era, 3.13, "they should be equal")
//
//}
//
//func TestStore_GetByPlayerId(t *testing.T) {
//	pitchFightInfos, err := pitcherFightStore.GetByPlayerId(playerId)
//	assert.Nil(t, err)
//	for _, v := range pitchFightInfos {
//		t.Logf("%+v", v)
//		assert.Equal(t, v.ID, uint(ID), "they should be equal")
//	}
//}
