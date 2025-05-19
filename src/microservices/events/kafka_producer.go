package main

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
    producer *kafka.Producer
}

func NewKafkaProducer(brokers string) *KafkaProducer {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": brokers,
    })
    if err != nil {
        panic(err)
    }
    return &KafkaProducer{producer: p}
}

func (kp *KafkaProducer) Produce(topic string, key string, value interface{}) error {
    // Реализация отправки сообщения
}
