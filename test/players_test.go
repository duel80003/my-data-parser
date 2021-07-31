package test

import (
	"cpbl-data-parser/datastore/players"
	"cpbl-data-parser/driver"
	"cpbl-data-parser/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

var store = players.New(driver.DatabaseClient())

func TestStore_BatchUpdate(t *testing.T) {
	player := &entities.Player{
		ID:   "test1234",
		Name: "test",
	}
	players := []*entities.Player{
		player,
	}

	store.BatchUpsert(players)
	assert.Equal(t, player.ID, "test1234", "they should be equal")
	assert.Equal(t, player.Name, "test", "they should be equal")
}

func TestStore_GetPlayers(t *testing.T) {
	players, err := store.GetPlayers()
	assert.Nil(t, err)
	length := len(players)

	assert.Greater(t, length, 0, "players should be greater than 0")
}
