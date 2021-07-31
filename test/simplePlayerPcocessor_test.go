package test

import (
	"my-data-parser/delivery"
	"testing"
)

func TestSimplePlayerProcessor(t *testing.T) {
	go delivery.AfterStop()
	delivery.SimplePlayerProcessor()
}
