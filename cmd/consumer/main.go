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
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

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
