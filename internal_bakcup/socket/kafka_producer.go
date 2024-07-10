package socket

import (
	"os"

	"github.com/IBM/sarama"
)

var Producer sarama.SyncProducer

func InitKafkaProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	var err error
	Producer, err = sarama.NewSyncProducer([]string{os.Getenv("Kafka_Producer_Host")}, config)
	if err != nil {
		panic(err)
	}
}
