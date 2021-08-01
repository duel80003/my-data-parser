package test

import (
	"github.com/stretchr/testify/assert"
	"my-data-parser/datastore/players"
	"my-data-parser/driver"
	"testing"
)

var store = players.New(driver.DatabaseClient())

//func TestStore_BatchUpdate(t *testing.T) {
//	player := &entities.Player{
//		ID:   "test1234",
//		Name: "test",
//	}
//	players := []*entities.Player{
//		player,
//	}
//
//	pitcherFightStore.BatchUpsert(players)
//	assert.Equal(t, player.ID, "test1234", "they should be equal")
//	assert.Equal(t, player.Name, "test", "they should be equal")
//}

func TestStore_GetPlayers(t *testing.T) {
	players, err := store.GetPlayers()
	assert.Nil(t, err)
	length := len(players)

	assert.Greater(t, length, 0, "players should be greater than 0")
}
