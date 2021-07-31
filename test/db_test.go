package test

import (
	"github.com/stretchr/testify/assert"
	"my-data-parser/driver"
	"testing"
)

func TestMySQLConnection(t *testing.T) {
	dbClient := driver.DatabaseClient()
	err := dbClient.Exec("select 1").Error
	assert.Nil(t, err, "database connection error", err)
}
