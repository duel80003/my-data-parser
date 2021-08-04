package main

import (
	"my-data-parser/delivery"
	"my-data-parser/utils"
	_ "my-data-parser/utils"
)

var logger = utils.LoggerInstance()

func main() {
	utils.SetLogLevel()
	ch := make(chan struct{})
	logger.Infof("service start.")
	delivery.Handler()
	<-ch
}
