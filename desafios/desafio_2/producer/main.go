package main

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer, err := ckafka.NewProducer(&ckafka.ConfigMap{"bootstrap.servers": "kafka:9092"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		panic(err)
	}

	msg := "Ola Consumer"
	topic := "desafioTopic"
	deliveryChan := make(chan ckafka.Event, 10000)

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	err = producer.Produce(message, deliveryChan)
	if err != nil {
		panic(err)
	}

	// err := p.Produce(&ckafka.Message{
	// 	TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
	// 	Value:          []byte(msg)
	// }, deliveryChan)
	DeliveryReport(deliveryChan)

}

func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}
