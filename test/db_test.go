package test

import (
	"cpbl-data-parser/driver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySQLConnection(t *testing.T) {
	dbClient := driver.DatabaseClient()
	err := dbClient.Exec("select 1").Error
	assert.Nil(t, err, "database connection error", err)
}
