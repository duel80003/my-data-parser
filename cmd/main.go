package main

import (
	"github.com/jasonlvhit/gocron"
	"my-data-parser/delivery"
	"my-data-parser/utils"
	_ "my-data-parser/utils"
)

var logger = utils.LoggerInstance()

func main() {
	utils.SetLogLevel()
	logger.Infof("service start.")
	err := gocron.Every(1).Day().At("15:30").Do(delivery.Handler)
	if err != nil {
		logger.Errorf("task error: %s", err)
	}
	<-gocron.Start()
}
