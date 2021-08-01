package test

import (
	"my-data-parser/delivery"
	"my-data-parser/utils"
	"testing"
)

func TestSimplePlayerProcessor(t *testing.T) {
	go delivery.AfterStop()
	delivery.SimplePlayerProcessor()
}

func TestStandingInfoProcessor(t *testing.T) {
	go delivery.AfterStop()
	delivery.StandingInfoProcessor()
}

func TestPlayerDetailsProcessor(t *testing.T) {
	utils.SetLogLevel()
	go delivery.AfterStop()
	delivery.PlayerDetailsProcessor()
}
