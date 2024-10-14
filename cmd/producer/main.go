package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	servers := []string{"localhost:9092"}
	topic := "topic_name"

	producer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			panic(err)
		}
	}(producer)

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("hello world"),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("partition: %d, offset: %d\n", partition, offset)
}
