package test

import (
	"my-data-parser/datastore/pitchingInfos"
	"my-data-parser/driver"
)

var pitchingStore = pitchingInfos.New(driver.DatabaseClient())

//const ID = 529
//const playerId = "test5678"
//
//func TestStore_BatchUpdate(t *testing.T) {
//	pitchingInfo := &entities.PitchingInfos{
//		PlayerId: playerId,
//		Era:      3.13,
//		Year:     "2021",
//		Acnt:     playerId,
//	}
//	pitchingInfo.SetID(ID)
//	pitcherFightInfos := []*entities.PitchingInfos{
//		pitchingInfo,
//	}
//
//	store.BatchUpsert(pitcherFightInfos)
//	assert.Equal(t, pitchingInfo.ID, uint(ID), "they should be equal")
//	assert.Equal(t, pitchingInfo.PlayerId, playerId, "they should be equal")
//	assert.Equal(t, pitchingInfo.Era, 3.13, "they should be equal")
//	assert.Equal(t, pitchingInfo.Year, "2021", "they should be equal")
//	assert.Equal(t, pitchingInfo.Acnt, playerId, "they should be equal")
//}
//
//func TestStore_GetByPlayerId(t *testing.T) {
//	pitchFightInfos, err := store.GetByPlayerId(playerId)
//	assert.Nil(t, err)
//	for _, v := range pitchFightInfos {
//		t.Logf("%+v", v)
//		assert.Equal(t, v.ID, uint(ID), "they should be equal")
//		assert.Equal(t, v.Year, "2021", "they should be equal")
//	}
//}
