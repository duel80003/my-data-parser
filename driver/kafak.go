package driver

import (
	"cpbl-data-parser/utils"
	"github.com/segmentio/kafka-go"
)

var (
	partition         = 0
	logger            = utils.LoggerInstance()
	kafkaHost         = utils.GetEnv("KAKFA_HOST")
	simplePlayerTopic = utils.GetEnv("SIMPLE_PLAYER_INFO_TOPIC")
	standingInfoTopic = utils.GetEnv("STANDING_INFO_TOPIC")
	playerDetailTopic = utils.GetEnv("PLAYER_DETAIL_TOPIC")
)

func newKafkaReader(topic string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaHost},
		Topic: topic,
		Partition: partition,
		MinBytes: 10e1,
		MaxBytes: 10e6,
		GroupID: topic,
	})
	return r
}

func SimplePlayerReader() *kafka.Reader {
	return newKafkaReader(simplePlayerTopic)
}

func StandingInfoReader() *kafka.Reader {
	return newKafkaReader(standingInfoTopic)
}

func PlayerDetailReader() *kafka.Reader {
	return newKafkaReader(playerDetailTopic)
}