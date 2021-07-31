package delivery

import (
	"context"
	"encoding/json"
	playerStore "my-data-parser/datastore/players"
	"my-data-parser/driver"
	"my-data-parser/entities"
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
)

func Handler() {
	utils.SetLogLevel()
	AfterStop()
	go SimplePlayerProcessor()
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
