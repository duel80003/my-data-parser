package use_cases

import (
	"fmt"
	"my-data-parser/entities"
	"my-data-parser/utils"
	"strconv"
)

var logger = utils.LoggerInstance()

func ParseToStandingInfos(data map[string]interface{}) (result []*entities.StandingInfos) {
	content, ok := data["content"]
	if !ok {
		return nil
	}

	array := content.([]interface{})

	for i, value := range array {
		standingInfo := &entities.StandingInfos{}
		tmpValue := value.([]interface{})
		standingInfo.SortNumber = i + 1
		standingInfo.TeamName = toString(tmpValue[0])
		standingInfo.Games = toInt(tmpValue[1])
		standingInfo.WinAndLose = toString(tmpValue[2])
		standingInfo.WinningPercentage = toFloat(tmpValue[3])
		standingInfo.GamesBehind = toString(tmpValue[4])
		standingInfo.GamesEliminate = toInt(tmpValue[5])
		standingInfo.Brothers = toString(tmpValue[6])
		standingInfo.ULions = toString(tmpValue[7])
		standingInfo.Monkeys = toString(tmpValue[8])
		standingInfo.Guardians = toString(tmpValue[9])
		standingInfo.Dragons = toString(tmpValue[10])
		standingInfo.RecordAtHome = toString(tmpValue[11])
		standingInfo.RecordAtAway = toString(tmpValue[12])
		standingInfo.CurrentStreak = toString(tmpValue[13])
		standingInfo.Last10Games = toString(tmpValue[14])
		logger.Infof("standingInfo %+v", standingInfo)
		result = append(result, standingInfo)
	}
	return result
}

func toString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func toInt(value interface{}) int {
	val, err := strconv.Atoi(toString(value))
	if err != nil {
		return 0
	}
	return val
}

func toFloat(value interface{}) float64 {
	val, err := strconv.ParseFloat(toString(value), 64)
	if err != nil {
		return 0
	}
	return val
}
