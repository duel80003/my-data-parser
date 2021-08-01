package use_cases

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"my-data-parser/entities"
	"regexp"
	"strings"
)

//func ParsePlayerDetailsData(rawData []byte) {
//	m := make(map[string]interface{})
//	_ = json.Unmarshal(rawData, &m)
//
//	dataMap := m["data"]
//	x := dataMap.(map[string]interface{})
//	logger.Infof("data %+v", dataMap.(map[string]interface{}))
//
//	playerInfoRawData := m["playerInfo"]
//
//	playerInfo := parsePlayerInfo(playerInfoRawData.(string))
//
//}

func base64Decode(s string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	return decoded, err
}

func ParsePlayerInfo(data interface{}) *entities.Player {
	logger.Infof("default ParsePlayerInfo data %v", data)
	if data == nil {
		return nil
	}
	decoded, err := base64Decode(data.(string))
	if err != nil {
		return nil
	}
	player := entities.Player{}
	_ = json.Unmarshal(decoded, &player)
	player.SetPlayerType()
	return &player
}

func ParseDefenceData(data interface{}, infos []*entities.DefenceInfos) (d []*entities.DefenceInfos) {
	logger.Infof("default ParseDefenceData data %v", data)
	if data == nil {
		return nil
	}
	decoded, err := base64Decode(data.(string))
	if err != nil {
		return nil
	}
	logger.Infof("ParseDefenceData decoded data %s", decoded)
	dataTable := entities.DataTable{}
	_ = json.Unmarshal(decoded, &dataTable)
	logger.Debugf("data table %+v", dataTable)
	content := fmt.Sprintf("%v", dataTable.TableContent)
	m := make(map[string][]*entities.DefenceInfos)
	var re1 = regexp.MustCompile(`[ Year{=\n"]+`)
	var re2 = regexp.MustCompile(` [ TeamNo=}\n]+`)
	_ = json.Unmarshal([]byte(content), &m)
	for k, val := range m {
		x := strings.Split(k, ",")
		x1 := re1.ReplaceAll([]byte(x[0]), []byte(""))
		x2 := re2.ReplaceAll([]byte(x[1]), []byte(""))
		for i := 0; i < len(val); i++ {
			val[i].Year = string(x1)
			val[i].TeamNo = string(x2)
			val[i].SetPlayerId(dataTable.ID)
			for _, v := range infos {
				if val[i].Year == v.Year && val[i].DefendStationName == v.DefendStationName {
					val[i].ID = v.ID
					break
				}
			}
		}
		d = append(d, val...)
	}
	return d
}

func ParsePitchingInfos(data interface{}, infos []*entities.PitchingInfos) (p []*entities.PitchingInfos) {
	logger.Infof("default ParsePitchingInfos data %v", data)
	if data == nil {
		return nil
	}
	decoded, err := base64Decode(data.(string))
	if err != nil {
		return nil
	}
	logger.Debugf("ParsePitchingInfos decoded data %s", decoded)
	dataTable := entities.DataTable{}
	_ = json.Unmarshal(decoded, &dataTable)
	content := fmt.Sprintf("%v", dataTable.TableContent)
	_ = json.Unmarshal([]byte(content), &p)
	for i := 0; i < len(p); i++ {
		for _, val := range infos {
			if p[i].Year == val.Year {
				p[i].ID = val.ID
				break
			}
		}
		p[i].SetPlayerId(dataTable.ID)
	}
	return p
}

func ParsePitcherFollowInfos(data interface{}, infos []*entities.PitcherFollowInfos) (p []*entities.PitcherFollowInfos) {
	logger.Infof("default ParsePitcherFollowInfos data %v", data)
	if data == nil {
		return nil
	}
	decoded, err := base64Decode(data.(string))
	if err != nil {
		return nil
	}
	logger.Debugf("parsePitcherFollowInfos decoded data %s", decoded)
	dataTable := entities.DataTable{}
	_ = json.Unmarshal(decoded, &dataTable)
	content := fmt.Sprintf("%v", dataTable.TableContent)
	_ = json.Unmarshal([]byte(content), &p)
	for i := 0; i < len(p); i++ {
		for _, val := range infos {
			if p[i].SId.Value == val.SId.Value {
				p[i].ID = val.ID
				break
			}
		}
		p[i].SetPlayerId(dataTable.ID)
	}
	return p
}

func ParsePitcherFightInfos(data interface{}, infos []*entities.PitcherFightInfos) (p []*entities.PitcherFightInfos) {
	logger.Infof("default ParsePitcherFightInfos data %v", data)
	if data == nil {
		return nil
	}
	decoded, err := base64Decode(data.(string))
	if err != nil {
		return nil
	}
	logger.Debugf("ParsePitcherFightInfos decoded data %s", decoded)
	dataTable := entities.DataTable{}
	_ = json.Unmarshal(decoded, &dataTable)
	content := fmt.Sprintf("%v", dataTable.TableContent)
	_ = json.Unmarshal([]byte(content), &p)
	for i := 0; i < len(p); i++ {
		for _, val := range infos {
			p[i].ID = val.ID
			break
		}
		p[i].SetPlayerId(dataTable.ID)
	}
	return p
}

func ParseBatting(data interface{}, originInfos []*entities.BattingInfos) (b []*entities.BattingInfos) {
	logger.Infof("default parseBatting data %v", data)
	if data == nil {
		return nil
	}
	decoded, err := base64Decode(data.(string))
	if err != nil {
		return nil
	}
	logger.Debugf("parseBatting decoded data %s", decoded)
	dataTable := entities.DataTable{}
	_ = json.Unmarshal(decoded, &dataTable)
	content := fmt.Sprintf("%v", dataTable.TableContent)
	_ = json.Unmarshal([]byte(content), &b)
	for i := 0; i < len(b); i++ {
		for _, val := range originInfos {
			if b[i].Year == val.Year {
				b[i].ID = val.ID
				break
			}
		}
		b[i].SetPlayerId(dataTable.ID)
	}
	return b
}

func ParseBatterFollowInfos(data interface{}, originInfos []*entities.BatterFollowInfos) (b []*entities.BatterFollowInfos) {
	logger.Infof("default parseBatterFollowInfos data %v", data)
	if data == nil {
		return nil
	}
	decoded, err := base64Decode(data.(string))
	if err != nil {
		return nil
	}
	logger.Debugf("parseBatting decoded data %s", decoded)
	dataTable := entities.DataTable{}
	_ = json.Unmarshal(decoded, &dataTable)
	content := fmt.Sprintf("%v", dataTable.TableContent)
	logger.Debugf("parseBatterFollowInfos content %s", content)
	_ = json.Unmarshal([]byte(content), &b)
	for i := 0; i < len(b); i++ {
		for _, val := range originInfos {
			if b[i].SId.Value == val.SId.Value {
				b[i].ID = val.ID
				break
			}
		}
		b[i].SetPlayerId(dataTable.ID)
	}
	return b
}

func ParseBatterFightInfos(data interface{}, infos []*entities.BatterFightInfos) (p []*entities.BatterFightInfos) {
	logger.Infof("default parseBatting data %v", data)
	if data == nil {
		return nil
	}
	decoded, err := base64Decode(data.(string))
	if err != nil {
		return nil
	}
	logger.Debugf("parseBatting decoded data %s", decoded)
	dataTable := entities.DataTable{}
	_ = json.Unmarshal(decoded, &dataTable)
	content := fmt.Sprintf("%v", dataTable.TableContent)
	logger.Debugf("parseBatterFightInfos content %s", content)
	_ = json.Unmarshal([]byte(content), &p)
	for i := 0; i < len(p); i++ {
		for _, val := range infos {
			p[i].ID = val.ID
			break
		}
		p[i].SetPlayerId(dataTable.ID)
	}
	return p
}
