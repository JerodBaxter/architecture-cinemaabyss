package main

type KafkaConsumer struct {
    consumer *kafka.Consumer
}

func NewKafkaConsumer(brokers string) *KafkaConsumer {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": brokers,
        "group.id":         "events-service",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    return &KafkaConsumer{consumer: c}
}

func (kc *KafkaConsumer) Consume(topics []string) {
    // Реализация потребления сообщений
}
