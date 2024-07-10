package socket

import (
	"fmt"

	"github.com/IBM/sarama"
)

var Consumer sarama.Consumer
var ConsumerGroup sarama.ConsumerGroup

func InitKafkaConsumer() {
	config := sarama.NewConfig()
	config.ChannelBufferSize = 1000000

	client, err := sarama.NewClient([]string{"localhost:29092"}, config)

	if err != nil {
		panic(err)
	}

	// if you want completed message -> 0
	lastoffset, err := client.GetOffset("eklee_topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("eklee_topic", 0, lastoffset)
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("Closing consumer")
		if err := partitionConsumer.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	consumed := 0
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Topic %s Consumed message offset %d , Partition %d\n", msg.Topic, msg.Offset, msg.Partition)
			consumed++
			fmt.Printf("Consumed: %d\n", consumed)
			fmt.Println(string(msg.Key))
			fmt.Println(string(msg.Value))
			fmt.Println("")
		}
	}
}
