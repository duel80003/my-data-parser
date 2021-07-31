package main

import (
	"cpbl-data-parser/delivery"
	"time"
)

func main() {
	delivery.Handler()
	time.Sleep(time.Second * 20)
}
