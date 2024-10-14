package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	servers := []string{"localhost:9092"}
	topic := "topic_name"

	consumer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer func(consumer sarama.Consumer) {
		err := consumer.Close()
		if err != nil {
			panic(err)
		}
	}(consumer)

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func(partitionConsumer sarama.PartitionConsumer) {
		err := partitionConsumer.Close()
		if err != nil {
			panic(err)
		}
	}(partitionConsumer)

	fmt.Println("Consumer started")
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Message on %s: %s\n", msg.Topic, msg.Value)
		}

	}
}
