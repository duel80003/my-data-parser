package test

import (
	"my-data-parser/datastore/defenceInfos"
	"my-data-parser/driver"
)

var defenceStore = defenceInfos.New(driver.DatabaseClient())

//const ID = 1714
//const playerId = "test1234"
//
//func TestStore_BatchUpdate(t *testing.T) {
//	defenceInfo := &entities.DefenceInfos{
//		PlayerId:          playerId,
//		Year:              "2021",
//		DefendStationName: "test name",
//	}
//	defenceInfo.SetID(ID)
//	batterInfos := []*entities.DefenceInfos{
//		defenceInfo,
//	}
//
//	defenceStore.BatchUpsert(batterInfos)
//	assert.Equal(t, defenceInfo.ID, uint(ID), "they should be equal")
//	assert.Equal(t, defenceInfo.PlayerId, playerId, "they should be equal")
//	assert.Equal(t, defenceInfo.Year, "2021", "they should be equal")
//	assert.Equal(t, defenceInfo.DefendStationName, "test name", "they should be equal")
//
//}
//
//func TestStore_GetByPlayerId(t *testing.T) {
//	batterInfos, err := defenceStore.GetByPlayerId(playerId)
//	assert.Nil(t, err)
//	for _, v := range batterInfos {
//		t.Logf("%+v", v)
//		assert.Equal(t, v.ID, uint(ID), "they should be equal")
//		assert.Equal(t, v.Year, "2021", "they should be equal")
//		assert.Equal(t, v.DefendStationName, "test name", "they should be equal")
//	}
//}
