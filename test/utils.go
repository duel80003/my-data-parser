package test

import (
	"my-data-parser/datastore/players"
	"my-data-parser/driver"
	"my-data-parser/entities"
)

var playerStore = players.New(driver.DatabaseClient())

var PlayerID = "test"

func createUser() {
	player := &entities.Player{
		Name: "test",
		ID:   PlayerID,
	}
	playerStore.Upsert(player)
}
