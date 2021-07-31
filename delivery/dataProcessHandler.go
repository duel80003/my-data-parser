package delivery

import (
	"context"
	"encoding/json"
	playerStore "my-data-parser/datastore/players"
	standingInfosStore "my-data-parser/datastore/standingInfos"

	"my-data-parser/driver"
	"my-data-parser/entities"
	use_cases "my-data-parser/use-cases"
	"my-data-parser/utils"
	"os"
	"os/signal"
	"syscall"
)

var (
	simplePlayerDataReader = driver.SimplePlayerReader()
	standingInfoReader     = driver.StandingInfoReader()
	playerDetailReader     = driver.PlayerDetailReader()
	logger                 = utils.LoggerInstance()
	dbClient               = driver.DatabaseClient()
	simplePlayerStore      = playerStore.New(dbClient)
	standingStore          = standingInfosStore.New(dbClient)
)

func Handler() {
	utils.SetLogLevel()
	AfterStop()
	go SimplePlayerProcessor()
	go StandingInfoProcessor()
}

func SimplePlayerProcessor() {
	logger.Infoln("SimplePlayerProcessor running")
	for {
		message, err := simplePlayerDataReader.ReadMessage(context.Background())
		if err != nil {
			logger.Errorf("Read simple player message error %s", err)
			break
		}
		player := &entities.Player{}
		_ = json.Unmarshal(message.Value, player)
		logger.Infof("message %+v", player)
		simplePlayerStore.Upsert(player)
	}
}

func StandingInfoProcessor() {
	logger.Infoln("StandingInfoProcessor running")
	for {
		message, err := standingInfoReader.ReadMessage(context.Background())
		if err != nil {
			logger.Errorf("Read standing info message error %s", err)
			break
		}
		t := make(map[string]interface{})
		_ = json.Unmarshal(message.Value, &t)
		logger.Infof("message %v", t)

		standsInfos := use_cases.ParseToStandingInfos(t)
		standingStore.BatchUpsert(standsInfos)
	}
}

func AfterStop() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		defer close(c)
		_, ok := <-c
		if ok {
			simplePlayerDataReader.Close()
			standingInfoReader.Close()
			playerDetailReader.Close()
			logger.Println("kafka connection closed")
			os.Exit(0)
		}
	}()
}
