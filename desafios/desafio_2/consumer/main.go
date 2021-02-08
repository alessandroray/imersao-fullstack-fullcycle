package main

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	topics := []string{"desafioTopic"}

	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Mensagem em %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()
}
