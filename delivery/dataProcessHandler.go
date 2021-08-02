package delivery

import (
	"context"
	"encoding/json"
	battingFightInfosStore "my-data-parser/datastore/batterFightInfos"
	battingFollowInfosStore "my-data-parser/datastore/batterFollowInfos"
	battingInfosStore "my-data-parser/datastore/battingInfos"
	defenceInfosStore "my-data-parser/datastore/defenceInfos"
	pitcherFightInfosStore "my-data-parser/datastore/pitcherFightInfos"
	pitcherFollowInfosStore "my-data-parser/datastore/pitcherFollowInfos"
	pitchingInfosStore "my-data-parser/datastore/pitchingInfos"
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
	playerInfoStore        = playerStore.New(dbClient)
	standingStore          = standingInfosStore.New(dbClient)
	defenceStore           = defenceInfosStore.New(dbClient)
	pitchingStore          = pitchingInfosStore.New(dbClient)
	pitcherFollowStore     = pitcherFollowInfosStore.New(dbClient)
	pitcherFightStore      = pitcherFightInfosStore.New(dbClient)
	battingStore           = battingInfosStore.New(dbClient)
	batterFollowStore      = battingFollowInfosStore.New(dbClient)
	batterFightStore       = battingFightInfosStore.New(dbClient)
)

func Handler() {
	AfterStop()
	go SimplePlayerProcessor()
	go StandingInfoProcessor()
	go PlayerDetailsProcessor()
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
		playerInfoStore.UpsertWithoutUpdate(player)
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
		infos := standingStore.GetAll()
		standsInfos := use_cases.ParseToStandingInfos(t, infos)
		standingStore.BatchUpsert(standsInfos)
	}
}

func PlayerDetailsProcessor() {
	logger.Infoln("PlayerDetailsProcessor running")
	for {
		message, err := playerDetailReader.ReadMessage(context.Background())
		if err != nil {
			logger.Errorf("Read player detail info message error %s", err)
			break
		}
		m := make(map[string]interface{})
		_ = json.Unmarshal(message.Value, &m)
		data := m["data"]
		dataMap := data.(map[string]interface{})
		logger.Debugf("raw data %+v", dataMap)
		playerInfoRawData := dataMap["playerInfo"]
		playerInfo := use_cases.ParsePlayerInfo(playerInfoRawData)
		logger.Infof("playerInfo %+v", playerInfo)
		if playerInfo == nil {
			continue
		}
		go playerInfoStore.Upsert(playerInfo)

		go func() {
			defenceInfosRawData := dataMap["defence"]
			defenceInfosInDB, _ := defenceStore.GetByPlayerId(playerInfo.ID)
			defenceInfos := use_cases.ParseDefenceData(defenceInfosRawData.(string), defenceInfosInDB)
			defenceStore.BatchUpsert(defenceInfos)
		}()

		if playerInfo.PlayerType == "pitcher" {
			go func() {
				pitchingInfosRawData := dataMap["pitch"]
				pitchingInfosInDB, _ := pitchingStore.GetByPlayerId(playerInfo.ID)
				pitchingInfos := use_cases.ParsePitchingInfos(pitchingInfosRawData, pitchingInfosInDB)
				pitchingStore.BatchUpsert(pitchingInfos)
			}()
			go func() {
				pitcherFollowInfosRawData := dataMap["follow"]
				pitcherFollowInfosInDB, _ := pitcherFollowStore.GetByPlayerId(playerInfo.ID)
				pitcherFollowInfos := use_cases.ParsePitcherFollowInfos(pitcherFollowInfosRawData, pitcherFollowInfosInDB)
				pitcherFollowStore.BatchUpsert(pitcherFollowInfos)
			}()
			go func() {
				pitcherFightInfosRawData := dataMap["fight"]
				pitcherFightInfosInDB, _ := pitcherFightStore.GetByPlayerId(playerInfo.ID)
				pitcherFightInfos := use_cases.ParsePitcherFightInfos(pitcherFightInfosRawData, pitcherFightInfosInDB)
				pitcherFightStore.BatchUpsert(pitcherFightInfos)
			}()
		} else {
			go func() {
				battingInfosRawData := dataMap["batting"]
				battingInfosInDB, _ := battingStore.GetByPlayerId(playerInfo.ID)
				battingInfos := use_cases.ParseBatting(battingInfosRawData, battingInfosInDB)
				battingStore.BatchUpsert(battingInfos)
			}()
			go func() {
				batterFollowInfosRawData := dataMap["follow"]
				batterFollowInfosInDB, _ := batterFollowStore.GetByPlayerId(playerInfo.ID)
				batterFollowsInfos := use_cases.ParseBatterFollowInfos(batterFollowInfosRawData, batterFollowInfosInDB)
				batterFollowStore.BatchUpsert(batterFollowsInfos)
			}()
			go func() {
				batterFightInfosRawData := dataMap["fight"]
				batterFightInfosInDB, _ := batterFightStore.GetByPlayerId(playerInfo.ID)
				batterFightInfos := use_cases.ParseBatterFightInfos(batterFightInfosRawData, batterFightInfosInDB)
				batterFightStore.BatchUpsert(batterFightInfos)
			}()
		}

	}
}

func AfterStop() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2, os.Interrupt)
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
